package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pcbfilters/types"
	"strings"

	"github.com/gorilla/mux"

	"unicode"
)

// Count articles stuff
func countArticlesManufacturer(
	articles []types.CpArticleComponent,
	manufacturer string,
	filtersApplied []types.IFiltersApplied,
	search string,
	price types.PriceRange, // Using a pointer for optional price
	onlyStock bool,
	storageType []string,
) int {
	count := 0
	values := make([]string, 0) // Initialize empty slice for filter values

	// Extract filter values if filters are applied
	if len(filtersApplied) > 0 {
		for _, filterGroup := range filtersApplied {
			if filterGroup.Category != "manufacturer" {
				for _, filter := range filterGroup.Filters {
					values = append(values, filter.Value)
				}
			}
		}
	}

	for _, article := range articles {
		if matchFilters(article, values, search, price, onlyStock, storageType, manufacturer) {
			count++
			// fmt.Println(article.Attributes.Name, count)
		}
	}

	return count
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func articleContainsString(article types.CpArticleComponent, values []string) bool {
	for _, data := range article.Attributes.Attributes {
		if containsString(values, data.Value) {
			return true
		}
	}
	return false
}

func matchFilters(article types.CpArticleComponent, values []string, search string, price types.PriceRange, onlyStock bool, storageType []string, manufacturer string) bool {
	// t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	// result, _, _ := transform.String(t, search)
	// lowerSearch := strings.ToLower(result)
	// fmt.Println(lowerSearch)

	// Combine all filter checks using short-circuit evaluation
	return (
	// Check filter attributes
	(len(values) == 0 || // No filters applied (always true)
		articleContainsString(article, values)) &&
		// Check manufacturer
		article.Attributes.Manufacturer == manufacturer &&
		// Check price (if provided)
		(article.Attributes.Price >= price.Min && article.Attributes.Price <= price.Max) &&
		// Check stock (if onlyStock is true)
		(!onlyStock || article.Attributes.Stock > 0) &&
		// Check storage type (if provided)
		(len(storageType) == 0 || containsString(storageType, article.Attributes.TypeStorage)))
	//&&
	// Check search terms (at least one field matches)
	// (lowerSearch == "" || // No search term
	// 	strings.Contains(lowerSearch, strings.ToLower(strings.RemoveAccents(article.Attributes.Name))) ||
	// 	strings.Contains(lowerSearch, strings.ToLower(strings.RemoveAccents(article.Attributes.Sku))) ||
	// 	strings.Contains(lowerSearch, strings.ToLower(strings.RemoveAccents(article.Attributes.Dne)))))
}

// Helper function to check if a string exists in a slice (case-insensitive)
func containsString(slice []string, str string) bool {
	for _, item := range slice {
		if strings.ToLower(item) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

// End count articles stuff

// Helper function to check if a slice contains a string
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func containsFromIndex(slice []string, str string, fromIndex int) bool {
	for i := fromIndex; i < len(slice); i++ {
		item := slice[i]
		if item == str {
			return true
		}
	}
	return false
}

func filterUniqueCheckBoxes(checkBoxes []types.FilterSectionItem, ids []string) []types.FilterSectionItem {
	var uniqueCheckBoxes []types.FilterSectionItem

	for checkIndex, checkBox := range checkBoxes {
		// fmt.Println("comparing", checkBox.Value)
		if !containsFromIndex(ids, checkBox.Value, checkIndex+1) {
			uniqueCheckBoxes = append(uniqueCheckBoxes, checkBox)
		}
	}

	return uniqueCheckBoxes
}

func maxSubsliceLength(data [][]types.FilterSectionItem) int {
	maxLength := 0
	for _, innerSlice := range data {
		if len(innerSlice) > maxLength {
			maxLength = len(innerSlice)
		}
	}
	return maxLength
}

// countarticleFilter stuff
// TODO: review this
// TODO: search
func findManufacturerFilter(filters []types.IFiltersApplied) (types.IFiltersApplied, bool) {
	for _, filter := range filters {
		if filter.Category == "manufacturer" {
			return filter, true
		}
	}
	return filters[0], false
}

func getManufacturerValues(filters []types.FilterSectionItem) []string {
	if filters == nil {
		return nil
	}
	var manufacturers []string
	for _, filter := range filters {
		manufacturers = append(manufacturers, filter.Value)
	}
	return manufacturers
}

// func accentRemove(s string) string {
// 	// Implementación de la función para remover acentos (depende de la librería usada)
// 	// Ejemplo usando la librería `iancoleman.me/go/locale"
// 	// return locale.ToLower(s, "UTF-8")
// 	// ...
// }

func countArticleaFilter(
	articles []types.CpArticleComponent,
	value types.Attribute,
	filtersApplied []types.IFiltersApplied,
	search string,
	price types.PriceRange, // Using a pointer for optional price
	onlyStock bool,
	storageType []string,
) int {
	manufacturerFilter, worked := findManufacturerFilter(filtersApplied)
	var manufacturers []string
	if worked {
		manufacturers = getManufacturerValues(manufacturerFilter.Filters)
	}
	// lowerSearch := search

	// fmt.Println(manufacturers[0], "==<")

	return len(manufacturers)
}

// Función para procesar los datos según los filtros
func setFilters(data []types.CpArticleComponent, paginateFilters []types.FilterSectionRow, componentFilters []types.ComponentFilter, filtersApplied []types.IFiltersApplied, search string, price types.PriceRange, onlyStock bool, storageType []string) ([]types.FilterSectionRow, error) {
	// ([]loadDataFromFile, error) {
	//TODO: Implementación de la lógica de filtrado
	var attributes []types.Attribute
	var attributesFiltered []types.Attribute
	var manufacturers []types.FilterSectionItem

	for _, article := range data {
		manufacturerIndex := -1

		for mIndex, manofacturer := range manufacturers {
			if manofacturer.Value == article.Attributes.Manufacturer {
				manufacturerIndex = mIndex
			}
		}

		if manufacturerIndex == -1 {
			countArticleManufacturer := countArticlesManufacturer(
				data,
				article.Attributes.Manufacturer,
				filtersApplied,
				search,
				price,
				onlyStock,
				storageType)

			//types.FilterSectionItem(emptystr, article.Attributes.Manufacturer, article.Attributes.Manufacturer, false, countArticleManufacturer, emptystr)
			var newManufacturer types.FilterSectionItem
			newManufacturer.GroupID = ""
			newManufacturer.Title = article.Attributes.Manufacturer
			newManufacturer.Value = article.Attributes.Manufacturer
			newManufacturer.Selected = false
			newManufacturer.Count = countArticleManufacturer
			newManufacturer.Category = ""

			manufacturers = append(manufacturers, newManufacturer)
		}

		for _, attribute := range article.Attributes.Attributes {
			// fmt.Println(attribute)
			attributes = append(attributes, attribute)
		}

		// fmt.Println(indexArticle, manufacturerIndex)
	}

	var ids []string
	// ids2 := make(map[string]types.Attribute)
	for _, attribute := range attributes {
		ids = append(ids, attribute.ID)
		// ids2[attribute.ID] = attribute
	}

	for attrIndex, attribute := range attributes {
		if containsFromIndex(ids, attribute.ID, attrIndex+1) {
			attributesFiltered = append(attributesFiltered, attribute)
		}
	}

	var filters []types.FilterSectionRow

	// fmt.Println(len(attributes), len(attributesFiltered))

	var newFilterRow types.FilterSectionRow
	newFilterRow.ID = "manufacturer"
	newFilterRow.Title = "Marca"
	newFilterRow.CheckBoxes = manufacturers
	filters = append(filters, newFilterRow)

	var mergedArrayFilters []types.FilterSectionItem
	if len(filtersApplied) > 0 {
		var temp [][]types.FilterSectionItem
		for _, category := range filtersApplied {
			temp = append(temp, category.Filters)
		}

		maxLenght := maxSubsliceLength(temp)
		for i := 0; i < maxLenght; i++ {
			for _, arr := range temp {
				if i < len(arr) {
					mergedArrayFilters = append(mergedArrayFilters, arr[i])
				}
			}
		}
	}

	for _, filter := range componentFilters {
		var checkBoxes []types.FilterSectionItem

		for _, attribute := range attributesFiltered {
			if filter.ID == attribute.ID {
				// TODO: finish this
				countFilters := countArticleaFilter(
					data,
					attribute,
					filtersApplied,
					search,
					price,
					onlyStock,
					storageType,
				)
				var newCheckbox types.FilterSectionItem
				newCheckbox.Title = attribute.Title
				newCheckbox.Value = attribute.Value
				newCheckbox.Selected = false
				newCheckbox.Category = filter.Title
				newCheckbox.GroupID = attribute.ID
				newCheckbox.Count = countFilters
				checkBoxes = append(checkBoxes, newCheckbox)
			}
		}
		ids = nil
		for _, checkbox := range checkBoxes {
			ids = append(ids, checkbox.Value)
		}

		filtered := filterUniqueCheckBoxes(checkBoxes, ids)

		var newFilter types.FilterSectionRow

		newFilter.ID = filter.ID
		newFilter.Title = filter.Title
		newFilter.CheckBoxes = filtered
		filters = append(filters, newFilter)

	}

	return filters, nil
}

// Manejador de la ruta API
func handleFilters(w http.ResponseWriter, r *http.Request) {
	var data types.IFilters
	// fmt.Println(r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Print("error decoding")
	}

	// var response error
	response, _ := setFilters(data.Data, data.PaginateFilters, data.ComponentFilters, data.FiltersApplied, data.Search, data.Price, data.OnlyStock, data.StorageType)

	// Codificar la respuesta en JSON
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/filters", handleFilters).Methods("POST")

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}
