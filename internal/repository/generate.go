package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i UserDBRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i UserCacheRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i LogRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i AccessDBRepository -o ./mocks/ -s "_minimock.go"
