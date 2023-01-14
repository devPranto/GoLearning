// w.Header().Set("Content-Type", "application/json")
// 	r.ParseMultipartForm(10 << 20)

// 	file, _, err := r.FormFile("file")
// 	if err != nil {
// 		fmt.Println("Error Retrieving the File")
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	tempFile, err := ioutil.TempFile("static", "uploadFile-*.png")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer tempFile.Close()
// 	fileBytes, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	tempFile.Write(fileBytes)
// 	person := user{

// 		Name:       r.FormValue("name"),
// 		City:       r.FormValue("city"),
// 		ProfilePic: "/" + path.Base(tempFile.Name()),
// 	}

// 	insertResult, err := userCollection.InsertOne(context.TODO(), person)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Inserted a single document: ", insertResult)
// 	json.NewEncoder(w).Encode(insertResult.InsertedID)