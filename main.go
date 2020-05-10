/*
 * Copyright 2020 Aletheia Ware LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var count = flag.Int("count", 10, "Fibonacci Number Count")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Usage: " + os.Args[0] + " <number>")
	} else {
		a := 0
		b, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < *count; i++ {
			log.Println(a)
			a, b = b, a+b
		}
	}
}
