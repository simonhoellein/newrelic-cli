package synthetics

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/newrelic/newrelic-cli/internal/client"
	"github.com/newrelic/newrelic-cli/internal/install/ux"
	"github.com/newrelic/newrelic-cli/internal/output"
	"github.com/newrelic/newrelic-cli/internal/utils"
	"github.com/newrelic/newrelic-client-go/v2/pkg/synthetics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	batchFile string
	guid      []string

	// Temporary variable to shift between scenarios of
	// mock JSON responses of the automatedTestResults query
	// scenario = 2

	// apiTimeout = time.Minute * 5

	// To be changed to 30 seconds in the real implementation, as suggested by the Synthetics team
	pollingInterval = time.Second * 30

	// Spinner
	//progressIndicator = ux.NewSpinnerProgressIndicator()
	progressIndicator = ux.NewSpinner()
)

var cmdRun = &cobra.Command{
	Use: "run",
	//TODO: Find the Precise description.
	Short:   "Interact with New Relic Synthetics batch monitors",
	Example: "newrelic synthetics run --help",
	Long:    "Interact with New Relic Synthetics batch monitors",
	Run: func(cmd *cobra.Command, args []string) {
		var testsBatchID string

		// Config holds values unmarshalled from the YAML file
		// var config StartAutomatedTestInput
		var config SyntheticsStartAutomatedTestInput
		if batchFile != "" || len(guid) != 0 {
			if batchFile != "" {
				// Unmarshal YAML file to get monitors and their properties
				// content, err := os.ReadFile(batchFile)
				content, err := os.ReadFile(batchFile)
				if err != nil {
					log.Fatal(err)
				}
				err = yaml.Unmarshal(content, &config)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(config)

				for _, test := range config.Tests {
					requestBody := fmt.Sprintf(`{"guid": "%s", "isBlocking": %v}`, string(test.MonitorGUID), test.Config.IsBlocking)
					fmt.Println(requestBody)

					// the following line may be no longer necessary (all references may be discarded)
					guid = append(guid, string(test.MonitorGUID))
				}

			} else if guid != nil {
				// this is yet to be tested
				var tests []synthetics.SyntheticsAutomatedTestMonitorInput
				for _, id := range guid {
					tests = append(tests, synthetics.SyntheticsAutomatedTestMonitorInput{
						MonitorGUID: synthetics.EntityGUID(id),
					})
				}

				config = SyntheticsStartAutomatedTestInput{
					Tests: tests,
				}

				log.Println(config)
			}

			// ----------------------------------------------------------------------------------
			// NOTE: mockBatchID is returned by the function that is (in reality) expected to
			// typecast inputs as needed by the API, send a request, and receive a batchID
			// as the response. Since this is being mocked, we only send the list of guids
			// to the function for now, but this is expected to belong to newrelic-client-go
			// and take all required parameters from the input YAML or GUIDs.
			// ----------------------------------------------------------------------------------

			testsBatchID = runSynthetics(guid, config)
			fmt.Println(testsBatchID)

		} else {
			utils.LogIfError(cmd.Help())
			//log.Fatal(" --batchFile <ymlFile> is required")
		}

		utils.LogIfFatal(output.Print(testsBatchID))

		// ----------------------------------------------------------------------------------
		// In order to mock implementation, the batchID has been hardcoded.
		// This is expected to be received in the response of syntheticsStartAutomatedTest.
		// ----------------------------------------------------------------------------------

		progressIndicator.Start("Fetching the status of tests in the batch....\n")

		// start := time.Now()

		// This variable has been used to track iterations of mock API calls for easier file opening of mock JSON files.
		// This may be discarded along with its usages after the API calling function is used.
		i := 0

		// An infinite loop
		for {

			// A timeout in the CLI may not be needed, based on recent suggestions received, as the API
			// returns a TIMED_OUT status if one or more job(s) in the batch consume > 10 minutes.
			// Update: Timeout is not needed.

			//if time.Since(start) > apiTimeout {
			//	fmt.Println("---------------------------")
			//	progressIndicator.Fail("Halting execution : reached timeout.")
			//	fmt.Println("---------------------------")
			//	break
			//}

			i++
			if i == 1 {
				// first iteration
				time.Sleep(time.Second * 30)
			}

			x := os.Getenv("NEW_RELIC_ACCOUNT_ID")
			log.Println()
			y, _ := strconv.Atoi(x)
			root, error := client.NRClient.Synthetics.GetAutomatedTestResult(y, testsBatchID)
			if error != nil {
				log.Fatal(error)
			}
			log.Println(root)
			// root := getAutomatedTestResultsMockResponse(testsBatchID, i)
			// ----------------------------------------------------------------------------------
			// ORIGINAL FUNCTION : Call the method from go-client that would send a request to the
			// automatedTestResults query with the batchID and fetch monitor details - sample below.

			// result, err := functionInClientGo(automatedTestResultQueryInput{batchID: batchID})
			// if err != nil {
			//	 return fmt.Errorf("Some error")
			// }
			// ----------------------------------------------------------------------------------

			exitStatus, ok := TestResultExitCodes[AutomatedTestResultsStatus(root.Status)]
			if !ok {
				handleStatus(*root, AutomatedTestResultsExitStatusUnknown)
			} else {
				handleStatus(*root, exitStatus)
			}

		}

	},
}

func init() {
	cmdRun.Flags().StringVarP(&batchFile, "batchFile", "b", "", "Input the YML file to batch and run the monitors")
	cmdRun.Flags().StringSliceVarP(&guid, "guid", "g", nil, "Batch the monitors using their guids and run the automated test")
	Command.AddCommand(cmdRun)

	// MarkFlagsMutuallyExclusive allows one flag at once be invoked
	cmdRun.MarkFlagsMutuallyExclusive("batchFile", "guid")

}

// getAutomatedTestResultsMockResponse is called to retrieve the mock JSON response of the automatedTestResults query
//func getAutomatedTestResultsMockResponse(batchID string, index int) (r AutomatedTestResult) {
//	directory := fmt.Sprintf("internal/synthetics/mock_json/Scenario %d", scenario)
//	filePath := fmt.Sprintf("%s/response_%d.json", directory, index)
//	jsonFile, err := os.Open(filePath)
//	if err != nil {
//		fmt.Println("Error opening file:", err)
//		return
//	}
//	defer func(jsonFile *os.File) {
//		err = jsonFile.Close()
//		if err != nil {
//			log.Fatal("Unable to close the file")
//		}
//	}(jsonFile)
//
//	byteValue, err := ioutil.ReadAll(jsonFile)
//	if err != nil {
//		fmt.Println("Error reading file:", err)
//		return
//	}
//
//	var root AutomatedTestResult
//	if err := json.Unmarshal(byteValue, &root); err != nil {
//		fmt.Println("Error unmarshalling JSON:", err)
//		return
//	}
//
//	return root
//}

// getMonitorTestsSummary is called every 15 seconds to print the status of individual monitors
func getMonitorTestsSummary(root synthetics.SyntheticsAutomatedTestResult) (string, string) {
	countSuccess := 0
	countFailure := 0
	countProgress := 0
	tests := root.Tests

	var summaryMessage string
	var resultMessage string

	for _, test := range tests {
		if test.Result == "SUCCESS" {
			countSuccess++
			resultMessage += fmt.Sprintf("\n - Success: %s (%s)", test.MonitorId, test.MonitorName)
		} else if test.Result == "FAILED" {
			countFailure++
			messageSubstring := ""
			if test.AutomatedTestMonitorConfig.IsBlocking == true {
				messageSubstring += fmt.Sprintf("(%s - Blocking)", test.MonitorName)
			} else if test.AutomatedTestMonitorConfig.IsBlocking == false {
				messageSubstring += fmt.Sprintf("(%s - Non-Blocking)", test.MonitorName)
			} else {
				messageSubstring += fmt.Sprintf("(%s)", test.MonitorName)
			}

			resultMessage += fmt.Sprintf("\n - Failed: %s %s", test.MonitorId, messageSubstring)
		} else if test.Result == "IN_PROGRESS" || test.Result == "" {
			countProgress++
			resultMessage += fmt.Sprintf("\n - In Progress: %s (%s)", test.MonitorId, test.MonitorName)
		}
	}

	summaryMessage = fmt.Sprintf("%d succeeded; %d failed; %d in progress.", countSuccess, countFailure, countProgress)
	return summaryMessage, resultMessage
}

// runSynthetics batches and call
func runSynthetics(guids []string, config SyntheticsStartAutomatedTestInput) string {
	fmt.Println("Running New Relic Synthetics for the following monitor GUID(s):")
	for _, guid := range guids {
		fmt.Println("-", guid)
	}

	// ----------------------------------------------------------------------------------
	// ORIGINAL FUNCTION : Call the method from go-client that would send a request to the
	// syntheticsStartAutomatedTest mutation (with the request body fit into datatypes
	// generated by Tutone) and would receive a response comprising a batchID.
	// ----------------------------------------------------------------------------------
	result, error := client.NRClient.Synthetics.SyntheticsStartAutomatedTest(config.Config, config.Tests)
	if error != nil {
		utils.LogIfFatal(error)
	}
	// returning a mock batchID.
	// Will be replaced with the response from the API
	return result.BatchId
}

// handleStatus processes the execution result of a test, taking into account
// the status contained in the root object and the associated exit status.
// Depending on the root.Status value, it prints an appropriate message, waits
// for the next API call, or exits the program with the given exit status.
//
// Parameters:
//   - root: Root struct that contains the status information.
//   - exitStatus: The AutomatedTestResultsExitStatus corresponding to the given root.Status.
//
// In the case of AutomatedTestResultsStatusInProgress, the function prints an
// information message, calls the getMonitorTestsSummary function, and waits for
// the specified pollingInterval before the next API call.
//
// In the cases of AutomatedTestResultsStatusTimedOut, AutomatedTestResultsStatusFailure,
// and AutomatedTestResultsStatusPassed, the function prints the execution result,
// calls the getMonitorTestsSummary function, and exits the program with the
// corresponding exit status code.
func handleStatus(root synthetics.SyntheticsAutomatedTestResult, exitStatus AutomatedTestResultsExitStatus) {
	retrievedStatus := string(root.Status)
	switch string(retrievedStatus) {
	case string(AutomatedTestResultsStatusInProgress):
		fmt.Println("\nStatus Received: IN_PROGRESS - re-calling the API in 15 seconds to fetch updated status...")
		summary, result := getMonitorTestsSummary(root)
		fmt.Printf("Summary: %s\n", summary)
		fmt.Println(result)
		fmt.Println()
		fmt.Println()
		time.Sleep(pollingInterval)
	case string(AutomatedTestResultsStatusTimedOut), string(AutomatedTestResultsStatusFailure), string(AutomatedTestResultsStatusPassed):
		progressIndicator.Success("Execution stopped - Status: " + retrievedStatus + "\n")
		fmt.Println("\nStatus Received: " + root.Status + " - Execution halted.")
		summary, result := getMonitorTestsSummary(root)
		fmt.Printf("Summary: %s\n", summary)
		fmt.Println(result)
		fmt.Println()
		fmt.Println()
		os.Exit(int(exitStatus))
	default:
		progressIndicator.Fail("Unexpected status: " + retrievedStatus)
		fmt.Println("\nStatus Received: " + root.Status + " - Exiting due to unexpected status.")
		os.Exit(int(AutomatedTestResultsExitStatusUnknown))
	}
}
