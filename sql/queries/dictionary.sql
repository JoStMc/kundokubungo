-- name: GetKanji :one
SELECT * FROM dictionary WHERE kanji = $1;
