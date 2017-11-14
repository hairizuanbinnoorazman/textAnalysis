# textAnalysis
API to handle a request to parse and analyze text.

The API will be able to accept a paragraph of text which it would then return a top array of 10 commonly found words in 
the text

## Quick overview

These are some of the steps done to handle the cases below:

Preprocessing:
- Remove punctuation
- Ensure that all letters are lowercase characters

Sorting (Several options here):
- (Current Implementation) Split into list of words array. Use hashmaps to collect the frequency count of each word and 
maintain a sorted list of the top 10 words in a separate slice.
   - Quick thoughts: Smaller array to loop through as compared to the potentially bigger spare array that the function might 
   take in. A very big assumption here is that it is assumed that most words are used sparingly, there are some words that 
   will be used very heavily compared to the rest. Most words would result in a count of 1.
- (Alternative Implementation) Split into list of words array. Construct the top 10 list by iterating through the array for each 
element in the top 10 list.
   - Quick thoughts: For each element, since the array is initally not sorted (hashmaps can't be sorted), that would mean 
   there would be need to compare every single element to find out the next element in line. 
- (Alternative Implementation) Split into list of words array. Add it to a hashed linked list which would internally sort itself
after each insertion. Sorting on insert is based of insertion sorts. Query for the top 10 is obtained by retrieving top 10
records from the list.
   - Quick thoughts: If given a huge block of text where most of the words would have very low frequency, that would mean that 
   of the sorting might not be too useful; unless one requests for a large list of such data. Also, imagine if you have 
   the word 'zoo' but infront of it, there are 10000 other words; that would mean that it might need to do 10000 
   comparisons before settling.

## Feature set

- Given a block of text to the API, return the top 10 words by default
- If there are less than 10 unique words, it will return with whatever that it has in the list
- The list of words being returned is sorted by frequency and then sorted alphatecially from a-z
- Words are returned as lowercase

## Running the service

Running the service without building the binary for test run purposes
```
# Go to the folder that contains server.go in the src/server folder
go run server.go

# Running a test run
curl http://localhost:3000/analysetext -H "Accept: application/json" -X POST -d '{"text":"a a a"}'
``` 

To build the service, ensure that the gopath and gobin are configured accordingly. This may differ between users.
Run either the commands
```
# Run either command:
go install server
go build server
```

