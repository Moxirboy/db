package postgres

const (
	GetQuestions = `
	SELECT id, questions_text 
	FROM questions
	ORDER BY RANDOM()
	LIMIT 50;

	`
	GetChoices = `
	select id,choice_text,is_correct from choices where question_id = $1
	`
	GetAnswer = `
	select choice_id from answer where question_id = $1 
	`
)