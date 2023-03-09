package postgres_test

// import (
// 	"database/sql/driver"
// 	"net/http"
// 	"testing"

// 	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
// 	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func TestAddStaff(t *testing.T) {
// 	// db, err := sql.Open("pgx", "uhuy")
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("connection error: $s", err)
// 	}
// 	defer db.Close()

// 	dialector := postgres.New(
// 		postgres.Config{
// 			Conn:                 db,
// 			PreferSimpleProtocol: true,
// 		},
// 	)

// 	gdb, err := gorm.Open(dialector, &gorm.Config{})

// 	payload := entity.Staff{
// 		StaffType: "staff",
// 		Name:      "Staff",
// 		Phone:     "08123456789",
// 		Email:     "staff@simars.com",
// 		Password:  "staff",
// 		Token:     "staff_token",
// 	}

// 	mock.ExpectBegin()
// 	mock.ExpectExec("INSERT INTO staffs").WithArgs(
// 		driver.Value("staff"),
// 		driver.Value("Staff"),
// 		driver.Value("08123456789"),
// 		driver.Value("staff@simars.com"),
// 		driver.Value("staff"),
// 		driver.Value("staff_token"),
// 	)
// 	mock.ExpectCommit()

// 	staffRepository := repository.NewStaffRepository(gdb)
// 	code, err := staffRepository.AddStaff(payload)

// 	assert := assert.New(t)
// 	assert.Equal(code, http.StatusOK, "should equal")
// 	assert.Nil(err)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
