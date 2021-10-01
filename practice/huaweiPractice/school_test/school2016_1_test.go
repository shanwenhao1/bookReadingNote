package school_test

import (
	"bookReadingNote/practice/huaweiPractice/school"
	"testing"
)

func TestRun(t *testing.T) {
	school.GradeRun()
}

func TestMain(m *testing.M) {
	m.Run()
}
