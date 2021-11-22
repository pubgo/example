package main

import (
	"fmt"
	"log"

	"github.com/hamba/avro"
)

// https://github.com/nrwiersma/avro-benchmarks

var Schema = `{
		"type": "record",
		"name": "simple",
		"namespace": "org.hamba.avro",
		"fields" : [
			{"name": "a", "type": "long"},
			{"name": "b", "type": "string"}
		]
	}`

type SimpleRecord struct {
	A int64  `avro:"a"`
	B string `avro:"b"`
}

func main() {
	schema := avro.MustParse(Schema)

	//in := SimpleRecord{A: 27, B: "foo"}
	in := map[string]interface{}{
		"a": int64(27), "b": "foo",
	}
	data, err := avro.Marshal(schema, in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)

	//out := SimpleRecord{}
	out := make(map[string]interface{})
	err = avro.Unmarshal(schema, data, &out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", out)

	//Output: [54 6 102 111 111]
	// {A:27 B:foo}
}
