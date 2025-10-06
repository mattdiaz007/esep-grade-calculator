package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if got := Assignment.String(); got != "assignment" {
		t.Fatalf("Assignment.String()=%q want %q", got, "assignment")
	}
	if got := Exam.String(); got != "exam" {
		t.Fatalf("Exam.String()=%q want %q", got, "exam")
	}
	if got := Essay.String(); got != "essay" {
		t.Fatalf("Essay.String()=%q want %q", got, "essay")
	}
}

func TestNoGrades_F(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("no grades: got %s want F", got)
	}
}

func TestGetGradeCAndD(t *testing.T) {
	t.Run("C_at_70", func(t *testing.T) {
		gc := NewGradeCalculator()
		gc.AddGrade("a", 70, Assignment)
		gc.AddGrade("e", 70, Exam)
		gc.AddGrade("s", 70, Essay)
		if got := gc.GetFinalGrade(); got != "C" {
			t.Fatalf("got %s want C", got)
		}
	})
	t.Run("D_at_60", func(t *testing.T) {
		gc := NewGradeCalculator()
		gc.AddGrade("a", 60, Assignment)
		gc.AddGrade("e", 60, Exam)
		gc.AddGrade("s", 60, Essay)
		if got := gc.GetFinalGrade(); got != "D" {
			t.Fatalf("got %s want D", got)
		}
	})
}
