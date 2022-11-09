package install

import (
	"github.com/newrelic/newrelic-cli/internal/diagnose"
	diagnose_mocks "github.com/newrelic/newrelic-cli/internal/diagnose/mocks"
	discovery_mocks "github.com/newrelic/newrelic-cli/internal/install/discovery/mocks"
	recipes_mocks "github.com/newrelic/newrelic-cli/internal/install/recipes/mocks"
	"testing"

	"github.com/newrelic/newrelic-cli/internal/install/discovery"
	"github.com/newrelic/newrelic-cli/internal/install/execution"
	"github.com/newrelic/newrelic-cli/internal/install/recipes"
	"github.com/newrelic/newrelic-cli/internal/install/types"
	"github.com/newrelic/newrelic-cli/internal/install/ux"
	"github.com/newrelic/newrelic-cli/internal/install/validation"
)

//type RecipeInstallBuilder struct {
//	//configValidator    diagnose_mocks.Validator
//	configValidator diagnose.Validator
//	recipeFetcher   *recipes.MockRecipeFetcher
//	//discoverer         discovery_mocks.Discoverer
//	discoverer         discovery.Discoverer
//	status             *execution.InstallStatus
//	mockOsValidator    *discovery.MockOsValidator
//	manifestValidator  *discovery.ManifestValidator
//	licenseKeyFetcher  *MockLicenseKeyFetcher
//	shouldInstallCore  func() bool
//	installerContext   types.InstallerContext
//	recipeLogForwarder *execution.MockRecipeLogForwarder
//	recipeVarProvider  *execution.MockRecipeVarProvider
//	recipeExecutor     *execution.MockRecipeExecutor
//	progressIndicator  *ux.SpinnerProgressIndicator
//	agentValidator     *validation.MockAgentValidator
//	recipeValidator    *validation.MockRecipeValidator
//	recipeDetector     *MockRecipeDetector
//	processes          []types.GenericProcess
//}

type RecipeInstallStub struct {
	testContext        *testing.T
	configValidator    diagnose.Validator
	recipeFetcher      recipes.RecipeFetcher
	discoverer         discovery.Discoverer
	status             *execution.InstallStatus //struct
	mockOsValidator    *discovery.Validator
	manifestValidator  *discovery.Validator
	licenseKeyFetcher  *LicenseKeyFetcher
	shouldInstallCore  func() bool
	installerContext   types.InstallerContext //struct
	recipeLogForwarder *execution.LogForwarder
	recipeVarProvider  *execution.MockRecipeVarProvider //needs interface
	recipeExecutor     *execution.RecipeExecutor
	progressIndicator  *ux.SpinnerProgressIndicator
	agentValidator     *validation.MockAgentValidator  //needs interface
	recipeValidator    *validation.MockRecipeValidator //needs interface
	recipeDetector     *MockRecipeDetector             //needs interface
	processes          []types.GenericProcess
}

func NewRecipeInstallWithStubs(t *testing.T) *RecipeInstallStub {
	recipeInstallWithStubs := &RecipeInstallStub{
		testContext:     t,
		configValidator: diagnose_mocks.NewValidator(t),
		discoverer:      discovery_mocks.NewDiscoverer(t),
		recipeFetcher:   recipes_mocks.NewRecipeFetcher(t),
	}

	return recipeInstallWithStubs
}

func (ris *RecipeInstallStub) WithConfigValidatorStub(stub *diagnose_mocks.Validator) *RecipeInstallStub {
	ris.configValidator = stub
	return ris
}

func (ris *RecipeInstallStub) WithDiscovererStub(stub *discovery_mocks.Discoverer) *RecipeInstallStub {
	ris.discoverer = stub
	return ris
}

func (ris *RecipeInstallStub) WithRecipeFetcher(stub *recipes_mocks.RecipeFetcher) *RecipeInstallStub {
	ris.recipeFetcher = stub
	return ris
}

func (ris *RecipeInstallStub) WithProgressIndicator(i *ux.SpinnerProgressIndicator) *RecipeInstallStub {
	ris.progressIndicator = i
	return ris
}

func (ris *RecipeInstallStub) Create() *RecipeInstall {
	recipeInstall := &RecipeInstall{}
	recipeInstall.configValidator = ris.configValidator
	recipeInstall.discoverer = ris.discoverer
	recipeInstall.recipeFetcher = ris.recipeFetcher

	return recipeInstall
}

//
//func NewRecipeInstallBuilder() *RecipeInstallBuilder {
//	rib := &RecipeInstallBuilder{
//		configValidator: old.NewMockConfigValidator(),
//		recipeFetcher:   recipes.NewMockRecipeFetcher(),
//		discoverer:      discovery.NewMockDiscoverer(),
//		processes:       []types.GenericProcess{},
//	}
//
//	statusReporter := execution.NewMockStatusReporter()
//	statusReporters := []execution.StatusSubscriber{statusReporter}
//	status := execution.NewInstallStatus(statusReporters, execution.NewPlatformLinkGenerator())
//	rib.status = status
//
//	rib.mockOsValidator = discovery.NewMockOsValidator()
//	rib.manifestValidator = discovery.NewMockManifestValidator(rib.mockOsValidator)
//	// Default to not skip core
//	rib.shouldInstallCore = func() bool { return true }
//	rib.installerContext = types.InstallerContext{}
//	rib.licenseKeyFetcher = NewMockLicenseKeyFetcher()
//	rib.recipeLogForwarder = execution.NewMockRecipeLogForwarder()
//	rib.recipeVarProvider = execution.NewMockRecipeVarProvider()
//	rib.recipeVarProvider.Vars = map[string]string{}
//	rib.recipeExecutor = execution.NewMockRecipeExecutor()
//	rib.progressIndicator = ux.NewSpinnerProgressIndicator()
//	rib.agentValidator = &validation.MockAgentValidator{}
//	rib.recipeValidator = &validation.MockRecipeValidator{}
//	rib.recipeDetector = &MockRecipeDetector{}
//
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithLibraryVersion(libraryVersion string) *RecipeInstallBuilder {
//	rib.recipeFetcher.LibraryVersion = libraryVersion
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithFetchRecipesVal(fetchRecipesVal []*types.OpenInstallationRecipe) *RecipeInstallBuilder {
//	rib.recipeFetcher.FetchRecipesVal = fetchRecipesVal
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithRecipeDetectionResult(detectionResult *recipes.RecipeDetectionResult) *RecipeInstallBuilder {
//	rib.recipeDetector.AddRecipeDetectionResult(detectionResult)
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithLicenseKeyFetchResult(result error) *RecipeInstallBuilder {
//	rib.licenseKeyFetcher.FetchLicenseKeyFunc = func(ctx context.Context) (string, error) {
//		return "", result
//	}
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithConfigValidatorError(err error) *RecipeInstallBuilder {
//	rib.configValidator.Error = err
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithDiscovererError(err error) *RecipeInstallBuilder {
//	rib.discoverer.Error = err
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithStatusReporter(statusReporter *execution.MockStatusSubscriber) *RecipeInstallBuilder {
//	statusReporters := []execution.StatusSubscriber{statusReporter}
//	status := execution.NewInstallStatus(statusReporters, execution.NewPlatformLinkGenerator())
//	rib.status = status
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithDiscovererValidatorError(err error) *RecipeInstallBuilder {
//	rib.mockOsValidator.Error = err
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) withShouldInstallCore(shouldSkipCore func() bool) *RecipeInstallBuilder {
//	rib.shouldInstallCore = shouldSkipCore
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithTargetRecipeName(name string) *RecipeInstallBuilder {
//	rib.installerContext.RecipeNames = append(rib.installerContext.RecipeNames, name)
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithRecipeExecutionError(err error) *RecipeInstallBuilder {
//	rib.recipeExecutor.ExecuteErr = err
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithOutput(value string) *RecipeInstallBuilder {
//	rib.recipeExecutor.SetOutput(value)
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithRecipeOutput(value []string) *RecipeInstallBuilder {
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithRecipeVarValues(vars map[string]string, err error) *RecipeInstallBuilder {
//	rib.recipeVarProvider.Vars = vars
//	rib.recipeVarProvider.Error = err
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithRecipeLogForwarder(optIn bool) *RecipeInstallBuilder {
//	rib.recipeLogForwarder.SetUserOptedIn(optIn)
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithProgressIndicator(i *ux.SpinnerProgressIndicator) *RecipeInstallBuilder {
//	rib.progressIndicator = i
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithAgentValidationError(e error) *RecipeInstallBuilder {
//	rib.agentValidator.Error = e
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithRecipeValidationError(e error) *RecipeInstallBuilder {
//	rib.recipeValidator.Error = e
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) WithRunningProcess(cmd string, name string) *RecipeInstallBuilder {
//	p := recipes.NewMockProcess(cmd, name, 0)
//	rib.processes = append(rib.processes, p)
//	return rib
//}
//
//func (rib *RecipeInstallBuilder) Build() *RecipeInstall {
//	recipeInstall := &RecipeInstall{}
//	recipeInstall.discoverer = rib.discoverer
//	recipeInstall.configValidator = rib.configValidator
//	recipeInstall.recipeFetcher = rib.recipeFetcher
//	recipeInstall.status = rib.status
//	recipeInstall.manifestValidator = rib.manifestValidator
//	recipeInstall.bundlerFactory = func(ctx context.Context, detections recipes.RecipeDetectionResults) RecipeBundler {
//		return recipes.NewBundler(ctx, detections)
//	}
//	recipeInstall.bundleInstallerFactory = func(ctx context.Context, manifest *types.DiscoveryManifest, recipeInstallerInterface RecipeInstaller, statusReporter StatusReporter) RecipeBundleInstaller {
//		return NewBundleInstaller(context.Background(), &types.DiscoveryManifest{}, recipeInstall, rib.status)
//	}
//	recipeInstall.shouldInstallCore = rib.shouldInstallCore
//	recipeInstall.InstallerContext = rib.installerContext
//	recipeInstall.licenseKeyFetcher = rib.licenseKeyFetcher
//	recipeInstall.recipeLogForwarder = rib.recipeLogForwarder
//	recipeInstall.recipeVarPreparer = rib.recipeVarProvider
//	recipeInstall.recipeExecutor = rib.recipeExecutor
//	recipeInstall.progressIndicator = rib.progressIndicator
//	recipeInstall.agentValidator = rib.agentValidator
//	recipeInstall.recipeValidator = rib.recipeValidator
//	recipeInstall.recipeDetectorFactory = func(ctx context.Context, repo *recipes.RecipeRepository) RecipeStatusDetector {
//		return rib.recipeDetector
//	}
//	mockProcessEvaluator := recipes.NewMockProcessEvaluator()
//	mockProcessEvaluator.WithProcesses(rib.processes)
//	recipeInstall.processEvaluator = mockProcessEvaluator
//
//	return recipeInstall
//}
