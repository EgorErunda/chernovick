package io

//Чтение байтов
//Интерфейс Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

//Эта функция проверяет, что буфер полностью заполнен перед тем, как вернуть значение. Если размер полученных данных отличается от размера буфера, то вы получите ошибку io.ErrUnexpectedEOF
func ReadFull(r Reader, buf []byte) (n int, err error)

//Чтобы прочесть 8 байт, достаточно сделать так:
// buf := make([]byte, 8)
// if _, err := io.ReadFull(r, buf); err != nil {
//     return err
// }

//Эта функция записывает доступные для чтения данные в ваш буфер, но не менее указанного количества байт
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)

type Writer interface {
	Write(p []byte) (n int, err error)
}

//Эта функция использует буфер в 32 КБ, чтобы прочитать из src и записать в dst. Если случится ошибка, отличная от io.EOF, копирование остановится и вернётся ошибка.
func Copy(dst Writer, src Reader) (written int64, err error)

//func CopuN() скопирует не больше указанного количества
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
