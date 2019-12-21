package test_test

import (
	"coca/core/domain/call_graph"
	"coca/core/models"
	"coca/core/support"
	"encoding/json"
	. "github.com/onsi/gomega"
	"testing"
)


func Test_should_generate_correct_files(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := call_graph.NewCallGraph()
	file := support.ReadFile("_fixtures/call_api_test.json")
	_ = json.Unmarshal(file, &parsedDeps)

	dotContent := analyser.Analysis("com.phodal.pholedge.book.BookController.createBook", *&parsedDeps)

	g.Expect(dotContent).To(Equal(`digraph G {
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookFactory.create";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getIsbn";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.command.CreateBookCommand.getName";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookRepository.save";
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.model.Book.getId";
"com.phodal.pholedge.book.BookController.createBook" -> "com.phodal.pholedge.book.BookService.createBook";
}
`))

}