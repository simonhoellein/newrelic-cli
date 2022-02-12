package recipes

import (
	"github.com/newrelic/newrelic-cli/internal/install/types"
)

var coreBundleRecipeNames = []string{
	types.InfraAgentRecipeName,
	types.LoggingRecipeName,
	types.GoldenRecipeName,
}

type Bundler struct {
	RecipeRepository *RecipeRepository
	RecipeDetector   *RecipeDetector
}

func NewBundler(rr *RecipeRepository) *Bundler {
	return newBundler(rr, NewRecipeDetector())
}

func newBundler(rr *RecipeRepository, rd *RecipeDetector) *Bundler {
	return &Bundler{
		RecipeRepository: rr,
		RecipeDetector:   rd,
	}
}

func (b *Bundler) CreateCoreBundle() *Bundle {
	var coreRecipes []*types.OpenInstallationRecipe
	for _, recipeName := range coreBundleRecipeNames {
		if r := b.RecipeRepository.FindRecipeByName(recipeName); r != nil {
			coreRecipes = append(coreRecipes, r)
		}
	}

	return b.CreateBundle(coreRecipes)
}

func (b *Bundler) CreateBundle(recipes []*types.OpenInstallationRecipe) *Bundle {

	bundle := &Bundle{}

	for _, r := range recipes {
		// recipe shouldn't have itself as dependency
		visited := map[string]bool{r.Name: true}
		bundle.AddRecipe(b.getBundleRecipeWithDependencies(r, visited))
	}

	// TODO: do detection here, and there
	//b.RecipeDetector.DetectRecipes()

	return bundle
}

func (b *Bundler) CreateBundleRecipe(recipe *types.OpenInstallationRecipe) *BundleRecipe {

	visited := map[string]bool{recipe.Name: true}
	return b.getBundleRecipeWithDependencies(recipe, visited)
}

func (b *Bundler) getBundleRecipeWithDependencies(recipe *types.OpenInstallationRecipe, visited map[string]bool) *BundleRecipe {

	bundleRecipe := &BundleRecipe{
		Recipe: recipe,
	}

	for _, d := range recipe.Dependencies {
		if !visited[d] {
			visited[d] = true
			if r := b.RecipeRepository.FindRecipeByName(d); r != nil {
				dr := b.getBundleRecipeWithDependencies(r, visited)
				bundleRecipe.Dependencies = append(bundleRecipe.Dependencies, dr)
			}
		}
	}

	return bundleRecipe
}

// Control status

//Recipe Candidate, recipe + collection of status
//Recipe context, capturing recipe intall info, timing, status..etc.

// func (b *Bundler) createAdditionalBundle(recipes []types.OpenInstallationRecipe) []types.OpenInstallationRecipe {
// 	_, a := createBundles(coreBundleRecipeNames, recipes)
// 	return a
// }
