package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Search existing word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStringsEqual(t, actual, expected)
	})

	t.Run("Search word not in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("missingword")

		assertErrorsEqual(t, err, missingValueError)
	})

}

func TestDictionaryAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("Add word/definitition pair", func(t *testing.T) {
		mockword := "testword"
		mockdefinition := "testvalue"
		dictionary.Add(mockword, mockdefinition)

		value, exists := dictionary.Search(mockword)

		assertStringsEqual(t, value, mockdefinition)
		assertErrorsEqual(t, exists, nil)
	})
}

func assertStringsEqual(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Actual %s not equal %s", actual, expected)
	}
}

func assertErrorsEqual(t *testing.T, actual, expected error) {
	t.Helper()

	if actual != expected {
		t.Errorf("Actual %q not equal %q", actual, expected)
	}
}
