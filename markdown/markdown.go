package markdown

import (
	"fmt"
	"strings"
)

/*
Heading         # h1
Paragraphs      use one or more blank line
Line Break      ending with more than 1 space
Bold	        **bold text**
Italic	        *italicized text*
Blockquote	    > blockquote    ->  <blockquote> </blockquote>
Unordered List	- First item
                - Second item
                - Third item
*/

// NOTE: embedding <em> into <strong> using ***TTT*** is not supported yet

func Translate(input string) string {
	allLines := strings.Split(input, "\n")
	if strings.HasSuffix(input, "\n") {
		allLines = allLines[0 : len(allLines)-1]
	}
	builder := &strings.Builder{}
	parseP(allLines, builder, 0, len(allLines)-1, 0)
	return builder.String()
}

func parseP(lines []string, builder *strings.Builder, fromLine, toLine, blockQuoteLevel int) {
	blockQuotePrefix := getQuotePrefix(blockQuoteLevel)

	thisPFromLine := fromLine
	processedToLine := -1
	for i := fromLine; i <= toLine; i++ {
		if lines[i] == blockQuotePrefix {
			builder.WriteString("<p>")
			parseQuote(lines, builder, thisPFromLine, i, blockQuoteLevel)
			builder.WriteString("</p>")

			thisPFromLine = i + 1
			processedToLine = i
		}
	}

	if processedToLine < 0 {
		parseQuote(lines, builder, fromLine, toLine, blockQuoteLevel)
	} else if processedToLine < toLine {
		parseQuote(lines, builder, processedToLine+1, toLine, blockQuoteLevel)
	}
}

func parseQuote(lines []string, builder *strings.Builder, fromLine, toLine, parentLevel int) {
	thisLevel := parentLevel + 1
	processedTo := fromLine - 1
	for i := fromLine; i <= toLine; i++ {
		line := lines[i]
		hasNextLine := i != toLine
		if lineInsideQuote(line, thisLevel) {
			if !hasNextLine || !lineInsideQuote(lines[i+1], thisLevel) {
				builder.WriteString("<blockquote>")
				parseP(lines, builder, processedTo+1, i, parentLevel+1)
				builder.WriteString("</blockquote>")
				processedTo = i
			}
		} else {
			if !hasNextLine || lineInsideQuote(lines[i+1], thisLevel) {
				parseList(lines, builder, processedTo+1, i, len(getQuotePrefix(parentLevel)))
				processedTo = i
			}
		}
	}
}

func lineInsideQuote(line string, level int) bool {
	if level <= 0 {
		return false
	}

	for i := level; i < level+10; i++ {
		if strings.HasPrefix(line, getQuotePrefix(i)) {
			return true
		}
	}

	return false
}

func getHeadPrefix(headLevel int) string {
	if headLevel == 0 {
		return ""
	}

	headPrefix := ""
	for i := 0; i < headLevel; i++ {
		headPrefix += "#"
	}
	headPrefix += " "
	return headPrefix
}

func getQuotePrefix(blockQuoteLevel int) string {
	if blockQuoteLevel <= 0 {
		return ""
	}

	blockQuotePrefix := ""
	for i := 0; i < blockQuoteLevel; i++ {
		blockQuotePrefix += ">"
	}
	blockQuotePrefix += " "
	return blockQuotePrefix
}

func parseList(lines []string, builder *strings.Builder, fromLine, toLine, skipFirst int) {
	processedTo := fromLine - 1
	listPrefix := "- "
	for i := fromLine; i <= toLine; i++ {
		thisLine := lines[i][skipFirst:]
		hasNextLine := i != toLine
		if strings.HasPrefix(thisLine, listPrefix) {
			if !hasNextLine || !strings.HasPrefix(lines[i+1][skipFirst:], listPrefix) {
				builder.WriteString("<ul>")
				for j := processedTo + 1; j <= i; j++ {
					builder.WriteString("<li>")
					parseH(lines, builder, j, skipFirst+len(listPrefix))
					builder.WriteString("</li>")
				}
				builder.WriteString("</ul>")
				processedTo = i
			}

		} else {
			if !hasNextLine || strings.HasPrefix(lines[i+1][skipFirst:], listPrefix) {
				for j := processedTo + 1; j <= i; j++ {
					parseH(lines, builder, j, skipFirst)
				}
				processedTo = i
			}
		}
	}
}

func parseH(lines []string, builder *strings.Builder, lineNumber, skipFirst int) {
	line := lines[lineNumber][skipFirst:]
	for i := 1; i <= 6; i++ {
		if strings.HasPrefix(line, getHeadPrefix(i)) {
			builder.WriteString(fmt.Sprintf("<h%d>", i))
			parseBr(lines, builder, lineNumber, skipFirst+len(getHeadPrefix(i)))
			builder.WriteString(fmt.Sprintf("</h%d>", i))
			return
		}
	}

	parseBr(lines, builder, lineNumber, skipFirst)
}

func parseBr(lines []string, builder *strings.Builder, lineNumber, skipFirst int) {
	line := lines[lineNumber][skipFirst:]
	parseBold(line, builder)
	if strings.HasSuffix(line, "  ") {
		builder.WriteString("<br/>")
	}
}

func parseBold(line string, builder *strings.Builder) {
	firstSignIndex := strings.Index(line, "**")
	if firstSignIndex < 0 || firstSignIndex >= len(line)-2 {
		parseItalic(line, builder)
		return
	}

	secondSignIndex := strings.Index(line[firstSignIndex+2:], "**")
	if secondSignIndex < 0 {
		parseItalic(line, builder)
		return
	}
	secondSignIndex = secondSignIndex + firstSignIndex + 2

	parseItalic(line[0:firstSignIndex], builder)
	builder.WriteString("<strong>")
	parseItalic(line[firstSignIndex+2:secondSignIndex], builder)
	builder.WriteString("</strong>")
	if secondSignIndex < len(line)-2 {
		parseBold(line[secondSignIndex+2:], builder)
	}
}

func parseItalic(line string, builder *strings.Builder) {
	firstSignIndex := strings.Index(line, "*")
	if firstSignIndex < 0 || firstSignIndex >= len(line)-1 {
		builder.WriteString(line)
		return
	}

	secondSignIndex := strings.Index(line[firstSignIndex+1:], "*")
	if secondSignIndex < 0 {
		builder.WriteString(line)
		return
	}
	secondSignIndex = secondSignIndex + firstSignIndex + 1

	builder.WriteString(line[0:firstSignIndex])
	builder.WriteString("<em>")
	builder.WriteString(line[firstSignIndex+1 : secondSignIndex])
	builder.WriteString("</em>")
	if secondSignIndex < len(line)-1 {
		parseItalic(line[secondSignIndex+1:], builder)
	}
}
