# 7.14 Пример: XML-декодирование на основе лексем

В разделе 4.5 было показано, как декодировать документы `JSON` в структуры Go с помощью функций `Marshal` и `Unmarshal`
из пакета `encoding/json`. Пакет `encoding/xml` предоставляет аналогичный API. Этот подход удобен, когда мы хотим
построить представление дерева документа, но для многих программ это не требуется. Пакет `encoding/xml`
также предоставляет для декодирования XML низкоуровневый API `на основе лексем`. При использовании стиля на основе
лексем синтаксический анализатор получает входные данные и генерирует поток лексем, главным образом четырех
видов (`StartElement`, `EndElement`, `CharData`, `Comment`), каждый из которых является конкретным типом
пакета `encoding/xml`. Каждый вызов (*xml.Decoder).Token возвращает лексему.

Ниже приведены существенные для нашего рассмотрения части API:

``` go
encoding/xml
package xml

type Name struct {
  Local string // Например, title или id
}

type Attr struct {
  Name Name
  Value string
}

// Token включает StartElement, EndElement, CharData и Comment,
// а также некоторые скрытые типы (не показаны).
type Token interface{}
type StartElement struct { // Например, <name>
  Name Name
  Attr []Attr
}

type EndElement struct { Name Name } // Например, </name>
type CharData []byte                 // Например, <p>CharData</p>
type Comment []byte                  // Например, <!-- Comment -->

type Decoder struct {/*...*/}
func NewDecoder(io.Reader) *Decoder
func (*Decoder) Token() (Token, error) // Возвращает очередную лексему
```

Не имеющий методов интерфейс `Token` также является примером `распознаваемого объединения (объединение типов)`.
Предназначение традиционного интерфейса наподобие `io.Reader` - сокрытие деталей конкретных типов, которые ему
соответствуют, так, чтобы могли быть созданы новые реализации. Все конкретные типы обрабатываются одинаково. Напротив,
набор конкретных типов, которые соответствуют распознаваемому объединению, является изначально фиксированным и не
скрытым. Типы распознаваемого объединения имеют несколько методов. Функции, которые работают с ними, выражаются в виде
набора `case` `type switch`, с различной логикой в каждом конкретном случае.

Приведенная далее программа `xmlselect` извлекает и выводит текст определенных элементов дерева XML-документа. С помощью
описанного выше API она может делать свою работу за один проход по входным данным, без построения дерева (см.
xmlselect.go):

``` go
import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // Стек имен элементов
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // Внесение в стек
		case xml.EndElement:
			stack = stack[:len(stack)-1] // Удаление из стека
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll указывает, содержит ли x элементы y в том же порядке.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
```

Каждый раз, когда цикл в `main` встречает `StartElement`, он помещает имя элемента в стек, а для каждого `EndElement`
удаляет имя из стека. API гарантирует, что лексемы `StartElement` и `EndElement` в последовательности будут корректно
соответствовать друг другу даже в неверно сформированных документах. Комментарии игнорируются. Когда
программа `xmlselect` обнаруживает `CharData`, она выводит текст, только если стек содержит все элементы с именами,
перечисленными в качестве аргументов командной строки в том же порядке.

Показанная ниже команда выводит текст всех элементов `h2`, которые находятся ниже двух уровней элементов div. Вводом
программы является спецификация XML, сама являющаяся XML-документом.

#### URL W3 из книги - 404. Актуальный https://www.w3.org/TR/xml11/

``` shell
$ .\fetch https://www.w3.org/TR/xml11/ | .\xmlselect.exe div div h2

html body div div h2: 1 Introduction
html body div div h2: 2 Documents
html body div div h2: 3 Logical Structures
html body div div h2: 4 Physical Structures
html body div div h2: 5 Conformance
html body div div h2: 6 Notation
html body div div h2: A References
html body div div h2: B Definitions for Character Normalization
html body div div h2: C Expansion of Entity and Character References (Non-Normative)
html body div div h2: D Deterministic Content Models (Non-Normative)
html body div div h2: E Autodetection of Character Encodings (Non-Normative)
html body div div h2: F W3C XML Working Group (Non-Normative)
html body div div h2: G W3C XML Core Working Group (Non-Normative)
html body div div h2: H Production Notes (Non-Normative)
html body div div h2: I Suggestions for XML Names (Non-Normative)
```

# Выводы

* Пакет `encoding/xml` в Go предоставляет API для работы с документами XML. Он позволяет декодировать документы XML в
  структуры Go и кодировать структуры Go обратно в XML. Также он предоставляет низкоуровневый API для декодирования XML
  на основе лексем;
* Стиль на основе лексем означает, что синтаксический анализатор получает входные данные и генерирует поток лексем.
  `Лексема` - это единица информации, которая генерируется синтаксическим анализатором. Каждый
  вызов `(*xml.Decoder).Token` возвращает одну лексему;
* Интерфейс Token является примером распознаваемого объединения. Это означает, что он позволяет работать с фиксированным
  набором типов, которые изначально определены и не скрыты;
* Типы распознаваемого объединения обрабатываются с помощью `type switch`, где каждый `case` имеет свою логику;
* API гарантирует, что лексемы `StartElement` и `EndElement` будут соответствовать друг другу даже в неверно
  сформированных документах.