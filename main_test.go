package main

import (
	//"errors"
	"errors"
	"reflect"
	"testing"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("Error while mocking")
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "role"}).AddRow(1, "Jane", "Jane@gmail.com", "Lead")
	TC:=[]struct{
		id int
		user *emp
		mockQuery   interface{}
		expectError error
		
	}{
		{
		id:1,
		user:&emp{1, "Jane", "Jane@gmail.com", "Lead" },
		mockQuery:   mock.ExpectQuery("SELECT * FROM employee WHERE id=?").WithArgs(1).WillReturnRows(rows),
		expectError: nil,
	},
	{
		id:          3,
		user:        nil,
		mockQuery:   mock.ExpectQuery("SELECT * FROM employee WHERE id=?").WithArgs(3).WillReturnError(errors.New("err")),
		expectError: errors.New("err"),
	},
	}
	
	for _, testCase := range TC{
		t.Run("", func(t *testing.T) {
		   user,err := GetById(db,testCase.id)
		   if err != nil && err.Error() != testCase.expectError.Error() {
			  t.Errorf("expected error: %v, got: %v", testCase.expectError, err)
		   }
		   if !reflect.DeepEqual(user, testCase.user) {
			  t.Errorf("expected user: %v, got: %v", testCase.user, user)
		   }
		})
	 }
	}
func TestRemoveById(t *testing.T) {
    db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
    if err != nil {
        t.Errorf("Error while mocking")
    }
    defer db.Close()

    TC := []struct {
        id          int
        mockQuery   interface{}
        expectError error
    }{
        {
            id: 3,
            mockQuery: []interface{}{
                mock.ExpectExec("DELETE FROM employee WHERE id=?").WithArgs(3).WillReturnResult(sqlmock.NewResult(3, 1)),
            },
            expectError: nil,
        },
        //fail
        {
            id: 1,
            mockQuery: []interface{}{
                mock.ExpectExec("DELETE FROM employee WHERE id=?").WithArgs(1).WillReturnError(errors.New("err")),
            },
                
            expectError: errors.New("err"),
        },
        
    }

    for _, tc := range TC {
        err := RemoveById(db,tc.id)
        fmt.Println(err)

        if err != nil && err.Error() != tc.expectError.Error() {
            t.Errorf("expected error: %v, got: %v", tc.expectError, err)
		}
    }
}
func TestUpdateById(t *testing.T){
	db,mock,err:=sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err!=nil{
		t.Errorf("Error while mocking")
	}
	defer db.Close()

	Tc:=[]struct{
		id int
		name string
		mockQuery []interface{}
		expectError error
	}{
		{
			//pass
			id:1,
			name:"Tarun",
			mockQuery: []interface{}{
				mock.ExpectExec("UPDATE employee SET name=? WHERE id=?").WithArgs("Tarun",1).WillReturnResult(sqlmock.NewResult(1,1)),	
			},
			expectError: nil,
		},
	
			
			//fail
			{
				id:7,
				name:"T",
				mockQuery:[]interface{}{
					mock.ExpectExec("UPDATE employee SET name=? WHERE id=?").WithArgs("T",7).WillReturnError(errors.New("err")),
				},
				expectError:errors.New("err"),
			},	
		}
		for _,tc := range Tc {
			err := UpdateById(db,tc.name,tc.id)
			if err != nil && err.Error() != tc.expectError.Error() {
				t.Errorf("expected error: %v, got: %v", tc.expectError, err)
			}
		}
	}


// func TestInsert(t *testing.T){
// 	db,mock,err:=sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	if err!=nil{
// 		t.Errorf("Error while mocking")
// 	}
// 	defer db.Close()

// 	Tc:=[]struct{
// 		id int
// 		name string
// 		email string
// 		role string
// 		mockQuery []interface{}
// 		expectError error
// 	}{
// 		{
// 			id:1,
// 			name:"Anurag",
// 			email:"anuragchaubey2@gmail.com",
// 			role:"Intern",
// 			mockQuery: []interface{}{
// 				mock.ExpectExec("INSERT INTO employee VALUES (?,?,?,?)").withArgs(1,"Anurag","anuragchaubey2@gmail.com","Intern").WillReturnResult(sqlmock.NewResult(1,1)),	
// 			},
// 			expectError: nil,
// 		},
			
// 			//fail
// 			{
// 				id:4,
// 				name:"Anurag",
// 			   email:"anuragchaubey2@gmail.com",
// 			   role:"Intern",
// 				mockQuery:[]interface{}{
// 					mock.ExpectExec("INSERT INTO employee VALUES(?,?,?,?)").withArgs(4,"Anurag","anuragchaubey2@gmail.com","Intern").WillReturnError(errors.New("err")),
// 				},
// 				expectError:errors.New("err"),
// 			},
// 		}

// 		for _,tc := range Tc {
// 			err := Insert(db,emp{tc.id,tc.name,tc.email,tc.role})
// 			fmt.Println(err)
	
// 			if err != nil && err.Error() != tc.expectError.Error() {
// 				t.Errorf("expected error: %v, got: %v", tc.expectError, err)
// 			}
// 		}
// 	}

	




// func TestUpdateById( t *testing.T) {

//     db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//     if err != nil {
//         t.Fatalf("an error'%s' was not expected when opening a stub database connection", err)
//     }
//     defer db.Close()
//     mock.ExpectExec("UPDATE employee WHERE id=?").withArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

// 	TC :=[]struct {
// 		id          int
// 		user        *emp
// 		mockQuery   interface{}
// 		expectError error
// 	}{
// 	{
// 		//if success

// 		id:          1,
// 		user:        &emp{1, "Tarun", "John@gmail.com", "Lead"},
// 		mockQuery:   mock.ExpectQuery("UPDATE employee WHERE id=?").WithArgs(1).WillReturnRows(rows),
// 		expectError: nil,
// 	},
// 	{
// 		//if error

// 		id:          3,
// 		user:        nil,
// 		mockQuery:   mock.ExpectQuery("UPDATE employee WHERE id=?").WithArgs(3).WillReturnError(errors.New("err")),
// 		expectError: errors.New("err"),
// 	},
// }
//     if err = UpdateById(db, Tc.user.name, 1); err != nil {
//         t.Errorf("error was not expected while updating stats: %s", err)
//     }
//     if err := mock.ExpectationsWereMet(); err != nil {
//         t.Errorf("there were unfulfilled expectations: %s", err )
//     }
// }
	

	
// func RemoveById(db *sql.DB,id int)(error){
    
//     q:="DELETE FROM employee WHERE id=?"
    
//     _,e:=db.Exec(q,id)
//     if e!=nil{
//         return e
//     }
//     //defer del.Close()
//     return nil
// }
