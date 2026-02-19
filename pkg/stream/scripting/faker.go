package scripting

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/dop251/goja"
)

// setupFaker creates a faker object with random data generation methods.
func (e *Engine) setupFaker(vm *goja.Runtime) {
	fakerObj := vm.NewObject()

	// Person
	_ = fakerObj.Set("name", func() string { return gofakeit.Name() })
	_ = fakerObj.Set("firstName", func() string { return gofakeit.FirstName() })
	_ = fakerObj.Set("lastName", func() string { return gofakeit.LastName() })
	_ = fakerObj.Set("email", func() string { return gofakeit.Email() })
	_ = fakerObj.Set("phone", func() string { return gofakeit.Phone() })
	_ = fakerObj.Set("username", func() string { return gofakeit.Username() })
	_ = fakerObj.Set("password", func(call goja.FunctionCall) goja.Value {
		lower := true
		upper := true
		numeric := true
		special := false
		length := 12
		if len(call.Arguments) >= 1 {
			length = int(call.Arguments[0].ToInteger())
		}
		return vm.ToValue(gofakeit.Password(lower, upper, numeric, special, false, length))
	})
	_ = fakerObj.Set("ssn", func() string { return gofakeit.SSN() })
	_ = fakerObj.Set("gender", func() string { return gofakeit.Gender() })

	// Address
	_ = fakerObj.Set("address", func() string { return gofakeit.Address().Address })
	_ = fakerObj.Set("street", func() string { return gofakeit.Street() })
	_ = fakerObj.Set("city", func() string { return gofakeit.City() })
	_ = fakerObj.Set("state", func() string { return gofakeit.State() })
	_ = fakerObj.Set("stateAbr", func() string { return gofakeit.StateAbr() })
	_ = fakerObj.Set("zipCode", func() string { return gofakeit.Zip() })
	_ = fakerObj.Set("country", func() string { return gofakeit.Country() })
	_ = fakerObj.Set("countryAbr", func() string { return gofakeit.CountryAbr() })
	_ = fakerObj.Set("latitude", func() float64 { return gofakeit.Latitude() })
	_ = fakerObj.Set("longitude", func() float64 { return gofakeit.Longitude() })

	// Internet
	_ = fakerObj.Set("url", func() string { return gofakeit.URL() })
	_ = fakerObj.Set("domainName", func() string { return gofakeit.DomainName() })
	_ = fakerObj.Set("domainSuffix", func() string { return gofakeit.DomainSuffix() })
	_ = fakerObj.Set("ipv4", func() string { return gofakeit.IPv4Address() })
	_ = fakerObj.Set("ipv6", func() string { return gofakeit.IPv6Address() })
	_ = fakerObj.Set("macAddress", func() string { return gofakeit.MacAddress() })
	_ = fakerObj.Set("userAgent", func() string { return gofakeit.UserAgent() })
	_ = fakerObj.Set("httpMethod", func() string { return gofakeit.HTTPMethod() })
	_ = fakerObj.Set("httpStatusCode", func() int { return gofakeit.HTTPStatusCode() })
	_ = fakerObj.Set("httpStatusCodeSimple", func() int { return gofakeit.HTTPStatusCodeSimple() })

	// Company
	_ = fakerObj.Set("company", func() string { return gofakeit.Company() })
	_ = fakerObj.Set("companySuffix", func() string { return gofakeit.CompanySuffix() })
	_ = fakerObj.Set("jobTitle", func() string { return gofakeit.JobTitle() })
	_ = fakerObj.Set("jobDescriptor", func() string { return gofakeit.JobDescriptor() })
	_ = fakerObj.Set("jobLevel", func() string { return gofakeit.JobLevel() })
	_ = fakerObj.Set("buzzWord", func() string { return gofakeit.BuzzWord() })
	_ = fakerObj.Set("bs", func() string { return gofakeit.BS() })

	// Finance
	_ = fakerObj.Set("creditCardNumber", func() string { return gofakeit.CreditCardNumber(nil) })
	_ = fakerObj.Set("creditCardType", func() string { return gofakeit.CreditCardType() })
	_ = fakerObj.Set("creditCardExp", func() string { return gofakeit.CreditCardExp() })
	_ = fakerObj.Set("creditCardCvv", func() string { return gofakeit.CreditCardCvv() })
	_ = fakerObj.Set("currency", func() string { return gofakeit.Currency().Short })
	_ = fakerObj.Set("currencyLong", func() string { return gofakeit.Currency().Long })
	_ = fakerObj.Set("price", func(call goja.FunctionCall) goja.Value {
		min := 1.0
		max := 1000.0
		if len(call.Arguments) >= 1 {
			min = call.Arguments[0].ToFloat()
		}
		if len(call.Arguments) >= 2 {
			max = call.Arguments[1].ToFloat()
		}
		return vm.ToValue(gofakeit.Price(min, max))
	})
	_ = fakerObj.Set("achAccount", func() string { return gofakeit.AchAccount() })
	_ = fakerObj.Set("achRouting", func() string { return gofakeit.AchRouting() })
	_ = fakerObj.Set("bitcoinAddress", func() string { return gofakeit.BitcoinAddress() })
	_ = fakerObj.Set("bitcoinPrivateKey", func() string { return gofakeit.BitcoinPrivateKey() })

	// Text/Lorem
	_ = fakerObj.Set("word", func() string { return gofakeit.Word() })
	_ = fakerObj.Set("sentence", func(call goja.FunctionCall) goja.Value {
		count := 5
		if len(call.Arguments) >= 1 {
			count = int(call.Arguments[0].ToInteger())
		}
		return vm.ToValue(gofakeit.Sentence(count))
	})
	_ = fakerObj.Set("paragraph", func(call goja.FunctionCall) goja.Value {
		count := 3
		sentenceCount := 5
		wordCount := 10
		if len(call.Arguments) >= 1 {
			count = int(call.Arguments[0].ToInteger())
		}
		return vm.ToValue(gofakeit.Paragraph(count, sentenceCount, wordCount, "\n"))
	})
	_ = fakerObj.Set("loremIpsumWord", func() string { return gofakeit.LoremIpsumWord() })
	_ = fakerObj.Set("loremIpsumSentence", func(call goja.FunctionCall) goja.Value {
		count := 5
		if len(call.Arguments) >= 1 {
			count = int(call.Arguments[0].ToInteger())
		}
		return vm.ToValue(gofakeit.LoremIpsumSentence(count))
	})
	_ = fakerObj.Set("question", func() string { return gofakeit.Question() })
	_ = fakerObj.Set("quote", func() string { return gofakeit.Quote() })
	_ = fakerObj.Set("phrase", func() string { return gofakeit.Phrase() })
	_ = fakerObj.Set("noun", func() string { return gofakeit.Noun() })
	_ = fakerObj.Set("verb", func() string { return gofakeit.Verb() })
	_ = fakerObj.Set("adverb", func() string { return gofakeit.Adverb() })
	_ = fakerObj.Set("adjective", func() string { return gofakeit.Adjective() })
	_ = fakerObj.Set("preposition", func() string { return gofakeit.Preposition() })

	// Date/Time
	_ = fakerObj.Set("date", func() string { return gofakeit.Date().Format("2006-01-02") })
	_ = fakerObj.Set("dateTime", func() string { return gofakeit.Date().Format("2006-01-02T15:04:05Z07:00") })
	_ = fakerObj.Set("futureDate", func() string { return gofakeit.FutureDate().Format("2006-01-02") })
	_ = fakerObj.Set("pastDate", func() string { return gofakeit.PastDate().Format("2006-01-02") })
	_ = fakerObj.Set("timeZone", func() string { return gofakeit.TimeZone() })
	_ = fakerObj.Set("timeZoneAbv", func() string { return gofakeit.TimeZoneAbv() })
	_ = fakerObj.Set("month", func() int { return gofakeit.Month() })
	_ = fakerObj.Set("monthString", func() string { return gofakeit.MonthString() })
	_ = fakerObj.Set("weekDay", func() string { return gofakeit.WeekDay() })
	_ = fakerObj.Set("year", func() int { return gofakeit.Year() })
	_ = fakerObj.Set("hour", func() int { return gofakeit.Hour() })
	_ = fakerObj.Set("minute", func() int { return gofakeit.Minute() })
	_ = fakerObj.Set("second", func() int { return gofakeit.Second() })
	_ = fakerObj.Set("nanosecond", func() int { return gofakeit.NanoSecond() })

	// Numbers
	_ = fakerObj.Set("int", func(call goja.FunctionCall) goja.Value {
		min, max := 0, 100
		if len(call.Arguments) >= 1 {
			min = int(call.Arguments[0].ToInteger())
		}
		if len(call.Arguments) >= 2 {
			max = int(call.Arguments[1].ToInteger())
		}
		if min >= max {
			return vm.ToValue(min)
		}
		return vm.ToValue(gofakeit.IntRange(min, max))
	})
	_ = fakerObj.Set("int8", func() int8 { return gofakeit.Int8() })
	_ = fakerObj.Set("int16", func() int16 { return gofakeit.Int16() })
	_ = fakerObj.Set("int32", func() int32 { return gofakeit.Int32() })
	_ = fakerObj.Set("int64", func() int64 { return gofakeit.Int64() })
	_ = fakerObj.Set("uint8", func() uint8 { return gofakeit.Uint8() })
	_ = fakerObj.Set("uint16", func() uint16 { return gofakeit.Uint16() })
	_ = fakerObj.Set("uint32", func() uint32 { return gofakeit.Uint32() })
	_ = fakerObj.Set("uint64", func() uint64 { return gofakeit.Uint64() })
	_ = fakerObj.Set("float", func(call goja.FunctionCall) goja.Value {
		min, max := 0.0, 100.0
		if len(call.Arguments) >= 1 {
			min = call.Arguments[0].ToFloat()
		}
		if len(call.Arguments) >= 2 {
			max = call.Arguments[1].ToFloat()
		}
		if min >= max {
			return vm.ToValue(min)
		}
		return vm.ToValue(gofakeit.Float64Range(min, max))
	})
	_ = fakerObj.Set("float32", func() float32 { return gofakeit.Float32() })
	_ = fakerObj.Set("float64", func() float64 { return gofakeit.Float64() })
	_ = fakerObj.Set("boolean", func() bool { return gofakeit.Bool() })
	_ = fakerObj.Set("digit", func() string { return gofakeit.Digit() })
	_ = fakerObj.Set("letter", func() string { return gofakeit.Letter() })
	_ = fakerObj.Set("hexColor", func() string { return gofakeit.HexColor() })
	_ = fakerObj.Set("rgbColor", func() []int { return gofakeit.RGBColor() })
	_ = fakerObj.Set("safeColor", func() string { return gofakeit.SafeColor() })

	// UUID
	_ = fakerObj.Set("uuid", func() string { return gofakeit.UUID() })

	// Hacker
	_ = fakerObj.Set("hackerPhrase", func() string { return gofakeit.HackerPhrase() })
	_ = fakerObj.Set("hackerAbbreviation", func() string { return gofakeit.HackerAbbreviation() })
	_ = fakerObj.Set("hackerAdjective", func() string { return gofakeit.HackerAdjective() })
	_ = fakerObj.Set("hackerNoun", func() string { return gofakeit.HackerNoun() })
	_ = fakerObj.Set("hackerVerb", func() string { return gofakeit.HackerVerb() })
	_ = fakerObj.Set("hackeringVerb", func() string { return gofakeit.HackeringVerb() })

	// App/Product
	_ = fakerObj.Set("appName", func() string { return gofakeit.AppName() })
	_ = fakerObj.Set("appVersion", func() string { return gofakeit.AppVersion() })
	_ = fakerObj.Set("appAuthor", func() string { return gofakeit.AppAuthor() })
	_ = fakerObj.Set("productName", func() string { return gofakeit.ProductName() })
	_ = fakerObj.Set("productCategory", func() string { return gofakeit.ProductCategory() })
	_ = fakerObj.Set("productDescription", func() string { return gofakeit.ProductDescription() })
	_ = fakerObj.Set("productFeature", func() string { return gofakeit.ProductFeature() })
	_ = fakerObj.Set("productMaterial", func() string { return gofakeit.ProductMaterial() })

	// Vehicle
	_ = fakerObj.Set("car", func() map[string]interface{} {
		v := gofakeit.Car()
		return map[string]interface{}{
			"type":         v.Type,
			"fuel":         v.Fuel,
			"transmission": v.Transmission,
			"brand":        v.Brand,
			"model":        v.Model,
			"year":         v.Year,
		}
	})
	_ = fakerObj.Set("carType", func() string { return gofakeit.CarType() })
	_ = fakerObj.Set("carMaker", func() string { return gofakeit.CarMaker() })
	_ = fakerObj.Set("carModel", func() string { return gofakeit.CarModel() })
	_ = fakerObj.Set("carFuelType", func() string { return gofakeit.CarFuelType() })
	_ = fakerObj.Set("carTransmissionType", func() string { return gofakeit.CarTransmissionType() })

	// Food
	_ = fakerObj.Set("fruit", func() string { return gofakeit.Fruit() })
	_ = fakerObj.Set("vegetable", func() string { return gofakeit.Vegetable() })
	_ = fakerObj.Set("breakfast", func() string { return gofakeit.Breakfast() })
	_ = fakerObj.Set("lunch", func() string { return gofakeit.Lunch() })
	_ = fakerObj.Set("dinner", func() string { return gofakeit.Dinner() })
	_ = fakerObj.Set("snack", func() string { return gofakeit.Snack() })
	_ = fakerObj.Set("dessert", func() string { return gofakeit.Dessert() })

	// Animal
	_ = fakerObj.Set("animal", func() string { return gofakeit.Animal() })
	_ = fakerObj.Set("animalType", func() string { return gofakeit.AnimalType() })
	_ = fakerObj.Set("petName", func() string { return gofakeit.PetName() })
	_ = fakerObj.Set("cat", func() string { return gofakeit.Cat() })
	_ = fakerObj.Set("dog", func() string { return gofakeit.Dog() })
	_ = fakerObj.Set("bird", func() string { return gofakeit.Bird() })

	// Emoji
	_ = fakerObj.Set("emoji", func() string { return gofakeit.Emoji() })
	_ = fakerObj.Set("emojiCategory", func() string { return gofakeit.EmojiCategory() })
	_ = fakerObj.Set("emojiAlias", func() string { return gofakeit.EmojiAlias() })
	_ = fakerObj.Set("emojiTag", func() string { return gofakeit.EmojiTag() })

	// Beer
	_ = fakerObj.Set("beerName", func() string { return gofakeit.BeerName() })
	_ = fakerObj.Set("beerStyle", func() string { return gofakeit.BeerStyle() })
	_ = fakerObj.Set("beerHop", func() string { return gofakeit.BeerHop() })
	_ = fakerObj.Set("beerYeast", func() string { return gofakeit.BeerYeast() })
	_ = fakerObj.Set("beerMalt", func() string { return gofakeit.BeerMalt() })
	_ = fakerObj.Set("beerIbu", func() string { return gofakeit.BeerIbu() })
	_ = fakerObj.Set("beerAlcohol", func() string { return gofakeit.BeerAlcohol() })
	_ = fakerObj.Set("beerBlg", func() string { return gofakeit.BeerBlg() })

	// Book
	_ = fakerObj.Set("bookTitle", func() string { return gofakeit.BookTitle() })
	_ = fakerObj.Set("bookAuthor", func() string { return gofakeit.BookAuthor() })
	_ = fakerObj.Set("bookGenre", func() string { return gofakeit.BookGenre() })

	// Movie
	_ = fakerObj.Set("movieName", func() string { return gofakeit.MovieName() })
	_ = fakerObj.Set("movieGenre", func() string { return gofakeit.MovieGenre() })

	// Game
	_ = fakerObj.Set("gamertag", func() string { return gofakeit.Gamertag() })

	// Celebrity
	_ = fakerObj.Set("celebrityActor", func() string { return gofakeit.CelebrityActor() })
	_ = fakerObj.Set("celebrityBusiness", func() string { return gofakeit.CelebrityBusiness() })
	_ = fakerObj.Set("celebritySport", func() string { return gofakeit.CelebritySport() })

	// File/MIME
	_ = fakerObj.Set("fileExtension", func() string { return gofakeit.FileExtension() })
	_ = fakerObj.Set("fileMimeType", func() string { return gofakeit.FileMimeType() })

	// Language
	_ = fakerObj.Set("language", func() string { return gofakeit.Language() })
	_ = fakerObj.Set("languageAbbreviation", func() string { return gofakeit.LanguageAbbreviation() })
	_ = fakerObj.Set("programmingLanguage", func() string { return gofakeit.ProgrammingLanguage() })

	// Pick from array
	_ = fakerObj.Set("pick", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return goja.Undefined()
		}
		arr := call.Arguments[0].Export()
		slice, ok := arr.([]interface{})
		if !ok {
			return goja.Undefined()
		}
		if len(slice) == 0 {
			return goja.Undefined()
		}
		idx := gofakeit.IntRange(0, len(slice)-1)
		return vm.ToValue(slice[idx])
	})

	// Shuffle array
	_ = fakerObj.Set("shuffle", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return vm.ToValue([]interface{}{})
		}
		arr := call.Arguments[0].Export()
		slice, ok := arr.([]interface{})
		if !ok {
			return vm.ToValue([]interface{}{})
		}
		result := make([]interface{}, len(slice))
		copy(result, slice)
		gofakeit.ShuffleAnySlice(result)
		return vm.ToValue(result)
	})

	// Sample N items from array
	_ = fakerObj.Set("sample", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return vm.ToValue([]interface{}{})
		}
		arr := call.Arguments[0].Export()
		n := int(call.Arguments[1].ToInteger())
		slice, ok := arr.([]interface{})
		if !ok || n <= 0 {
			return vm.ToValue([]interface{}{})
		}
		if n > len(slice) {
			n = len(slice)
		}
		result := make([]interface{}, len(slice))
		copy(result, slice)
		gofakeit.ShuffleAnySlice(result)
		return vm.ToValue(result[:n])
	})

	_ = vm.Set("faker", fakerObj)
}
