package solidity

import (
	"regexp"
	"strings"

	"github.com/silasdev78/silas-code-inspector/internal/domain"
	solpatterns "github.com/silasdev78/silas-code-inspector/internal/knowledge/solidity"
)

type Scanner struct {
	patterns []domain.Pattern
}

func NewScanner() *Scanner {
	return &Scanner{patterns: solpatterns.Patterns()}
}

func (s *Scanner) Scan(source string) []domain.Issue {
	cleanSource := removeComments(source)
	lines := strings.Split(cleanSource, "\n")
	var issues []domain.Issue

	for _, pattern := range s.patterns {
		re, err := regexp.Compile(pattern.Regex)
		if err != nil {
			continue
		}
		matches := re.FindAllStringIndex(cleanSource, -1)
		for _, match := range matches {
			lineNum := lineNumber(cleanSource, match[0])
			snippet := ""
			if lineNum > 0 && lineNum <= len(lines) {
				snippet = strings.TrimSpace(lines[lineNum-1])
			}

			if pattern.ID == "SOL-008" && isDelegatecallToSelf(lines, lineNum) {
				continue
			}

			if pattern.ID == "SOL-006" && isCEIPatternFollowed(lines, lineNum) {
				continue
			}

			if pattern.ID == "SOL-004" && hasInitializerModifier(lines) {
				continue
			}

			if pattern.ID == "SOL-024" && isWeakRandomnessSafe(lines, lineNum) {
				continue
			}

			if (pattern.ID == "SOL-020" || pattern.ID == "SOL-029") && isViewOrPureFunction(lines, lineNum) {
				continue
			}

			issues = append(issues, domain.Issue{
				Title:          pattern.Title,
				Description:    pattern.Description,
				Severity:       pattern.Severity,
				Line:           lineNum,
				Column:         match[0],
				Snippet:        snippet,
				Recommendation: pattern.Recommendation,
			})
		}
	}
	return issues
}

func removeComments(source string) string {
	block := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	out := block.ReplaceAllString(source, "")
	line := regexp.MustCompile(`//.*`)
	out = line.ReplaceAllString(out, "")
	natspec := regexp.MustCompile(`///.*`)
	out = natspec.ReplaceAllString(out, "")
	return out
}

func lineNumber(source string, offset int) int {
	return strings.Count(source[:offset], "\n") + 1
}

func isDelegatecallToSelf(lines []string, lineNum int) bool {
	if lineNum < 1 || lineNum > len(lines) {
		return false
	}
	lineContent := lines[lineNum-1]
	return strings.Contains(lineContent, "address(this)") && strings.Contains(lineContent, "delegatecall")
}

func isCEIPatternFollowed(lines []string, callLine int) bool {
	if callLine < 2 {
		return false
	}
	start := callLine - 6
	if start < 0 {
		start = 0
	}
	for i := start; i < callLine-1; i++ {
		line := strings.TrimSpace(lines[i])
		if strings.Contains(line, "=") && !strings.HasPrefix(line, "require") && !strings.HasPrefix(line, "assert") {
			return true
		}
	}
	return false
}

func hasInitializerModifier(lines []string) bool {
	for _, line := range lines {
		if strings.Contains(line, "initializer") || strings.Contains(line, "onlyImpl") {
			return true
		}
	}
	return false
}

func isWeakRandomnessSafe(lines []string, lineNum int) bool {
	if lineNum < 1 || lineNum > len(lines) {
		return false
	}
	line := strings.TrimSpace(lines[lineNum-1])
	if strings.Contains(line, "emit ") || strings.HasPrefix(line, "//") {
		return true
	}
	for i := lineNum - 3; i < lineNum; i++ {
		if i >= 0 && i < len(lines) {
			if strings.Contains(lines[i], "view ") || strings.Contains(lines[i], "pure ") {
				return true
			}
		}
	}
	return false
}

func isViewOrPureFunction(lines []string, lineNum int) bool {
	if lineNum < 1 || lineNum > len(lines) {
		return false
	}
	start := lineNum - 3
	if start < 0 {
		start = 0
	}
	end := lineNum + 2
	if end > len(lines) {
		end = len(lines)
	}
	for i := start; i < end; i++ {
		line := strings.TrimSpace(lines[i])
		if strings.Contains(line, "view ") || strings.Contains(line, "pure ") ||
			strings.Contains(line, "view\t") || strings.Contains(line, "pure\t") {
			return true
		}
	}
	return false
}
