package vksdk

import "encoding/json"

// DatabaseGetChairsResponse struct
type DatabaseGetChairsResponse struct{}

// TODO database.getChairs Returns list of chairs on a specified faculty.

// DatabaseGetCitiesResponse struct
type DatabaseGetCitiesResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Region string `json:"region"`
	} `json:"items"`
}

// DatabaseGetCities returns a list of cities
func (vk *VK) DatabaseGetCities(params map[string]string) (DatabaseGetCitiesResponse, error) {
	var response DatabaseGetCitiesResponse

	// TODO testing dataxase.getCities

	rawResponse, err := vk.Request("database.getCities", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// DatabaseGetCitiesByIDResponse struct
type DatabaseGetCitiesByIDResponse struct{}

// TODO database.getCitiesById Returns information about cities by their IDs.

// DatabaseGetCountriesResponse struct
type DatabaseGetCountriesResponse struct{}

// TODO database.getCountries Returns a list of countries.

// DatabaseGetCountriesByIDResponse struct
type DatabaseGetCountriesByIDResponse struct{}

// TODO database.getCountriesById Returns information about countries by their IDs.

// DatabaseGetFacultiesResponse struct
type DatabaseGetFacultiesResponse struct{}

// TODO database.getFaculties Returns a list of faculties (i.e., university departments).

// DatabaseGetRegionsResponse struct
type DatabaseGetRegionsResponse struct{}

// TODO database.getRegions Returns a list of regions.

// DatabaseGetSchoolClassesResponse struct
type DatabaseGetSchoolClassesResponse struct{}

// TODO database.getSchoolClasses Returns a list of school classes specified for the country.

// DatabaseGetSchoolsResponse struct
type DatabaseGetSchoolsResponse struct{}

// TODO database.getSchools Returns a list of schools.

// DatabaseGetUniversitiesResponse struct
type DatabaseGetUniversitiesResponse struct{}

// TODO database.getUniversities Returns a list of higher education institutions.
