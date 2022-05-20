package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/mock"
	"github.com/tech-thinker/go-cookiecutter/app/models"
	"github.com/tech-thinker/go-cookiecutter/app/repository/mocks"
)

func Test_todoRepo_Save(t *testing.T) {
	type fields struct {
		mockDB *mocks.OrmerMock
	}
	type args struct {
		ctx context.Context
		doc *models.Todo
	}

	tests := []struct {
		name    string
		fields  fields
		prepare func(*fields)
		args    args
		wantErr bool
	}{
		{
			name:   "with_valid_data_should_success",
			fields: fields{},
			prepare: func(f *fields) {
				f.mockDB.On("Insert", mock.Anything).Return(int64(1), nil)
			},
			args: args{ctx: context.Background(), doc: &models.Todo{Task: func() *string {
				s := "test"
				return &s
			}(), Done: false}},
			wantErr: false,
		},
		{
			name:   "with_invalid_data_should_fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.mockDB.On("Insert", mock.Anything).Return(int64(0), errors.New("error"))
			},
			args: args{ctx: context.Background(), doc: &models.Todo{Task: func() *string {
				s := "test"
				return &s
			}(), Done: false}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Preparing mocks
			tt.fields = fields{
				mockDB: mocks.NewOrmerMock(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			SUT := NewTodoRepo(
				tt.fields.mockDB,
			)

			err := SUT.Save(tt.args.ctx, tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoRepo.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoRepo_FindOne(t *testing.T) {
	type fields struct {
		mockDB *mocks.OrmerMock
		mockQS *mocks.QuerySeterMock
	}
	type args struct {
		ctx context.Context
		doc models.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(*fields)
		args    args
		want    models.Todo
		wantErr bool
	}{
		{
			name:   "with_valid_id_should_success",
			fields: fields{},
			prepare: func(f *fields) {
				f.mockDB.On("QueryTable", mock.Anything).Return(f.mockQS)
				f.mockQS.On("Filter", mock.Anything, mock.Anything).Return(f.mockQS)
				f.mockQS.On("One", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
					arg := args.Get(0).(*models.Todo)
					arg.ID = 1
					arg.Task = func() *string {
						s := "test"
						return &s
					}()
					arg.Done = false
				})

			},
			args: args{
				ctx: context.Background(),
				doc: models.Todo{
					Base: models.Base{
						ID: 1,
					},
				},
			},
			want: models.Todo{
				Base: models.Base{
					ID: 1,
				},
				Task: func() *string {
					s := "test"
					return &s
				}(),
				Done: false,
			},
			wantErr: false,
		},
		{
			name:   "with_invalid_id_should_fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.mockDB.On("QueryTable", mock.Anything).Return(f.mockQS)
				f.mockQS.On("Filter", mock.Anything, mock.Anything).Return(f.mockQS)
				f.mockQS.On("One", mock.Anything).Return(orm.ErrNoRows)
			},
			args: args{
				ctx: context.Background(),
				doc: models.Todo{
					Base: models.Base{
						ID: 1,
					},
				},
			},
			want:    models.Todo{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields = fields{
				mockDB: mocks.NewOrmerMock(t),
				mockQS: mocks.NewQuerySeterMock(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			SUT := NewTodoRepo(
				tt.fields.mockDB,
			)

			got, err := SUT.FindOne(tt.args.ctx, tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoRepo.FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoRepo.FindOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoRepo_Update(t *testing.T) {
	type fields struct {
		mockDB *mocks.OrmerMock
	}

	type args struct {
		ctx            context.Context
		doc            *models.Todo
		fieldsToUpdate []string
	}

	tests := []struct {
		name    string
		fields  fields
		prepare func(*fields)
		args    args
		wantErr bool
	}{
		{
			name:   "with_valid_data_should_success",
			fields: fields{},
			prepare: func(f *fields) {
				f.mockDB.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)
			},
			args: args{
				ctx: context.Background(),
				doc: &models.Todo{
					Base: models.Base{
						ID: 1,
					},
					Task: func() *string {
						s := fake.Sentence()
						return &s
					}(),
					Done: false,
				},
				fieldsToUpdate: []string{"Task", "Done"},
			},
			wantErr: false,
		},
		{
			name:   "with_invalid_data_should_fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.mockDB.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(int64(0), errors.New("error"))
			},
			args: args{
				ctx: context.Background(),
				doc: &models.Todo{
					Base: models.Base{
						ID: 1,
					},
					Task: func() *string {
						s := fake.Sentence()
						return &s
					}(),
					Done: false,
				},
				fieldsToUpdate: []string{"Task", "Done"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields = fields{
				mockDB: mocks.NewOrmerMock(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			SUT := NewTodoRepo(
				tt.fields.mockDB,
			)

			if err := SUT.Update(tt.args.ctx, tt.args.doc, tt.args.fieldsToUpdate); (err != nil) != tt.wantErr {
				t.Errorf("todoRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoRepo_FindAll(t *testing.T) {
	type fields struct {
		mockDB *mocks.OrmerMock
		mockQS *mocks.QuerySeterMock
	}
	type args struct {
		ctx   context.Context
		query models.TodoQuery
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(*fields)
		args    args
		want    []models.Todo
		want1   int64
		wantErr bool
	}{
		{
			name:   "with_valid_data_should_success",
			fields: fields{},
			prepare: func(f *fields) {
				f.mockDB.On("QueryTable", mock.Anything).Return(f.mockQS)
				f.mockQS.On("OrderBy", mock.Anything).Return(f.mockQS)
				f.mockQS.On("Filter", mock.Anything, mock.Anything).Return(f.mockQS)
				f.mockQS.On("All", mock.Anything).Return(int64(1), nil).Run(func(args mock.Arguments) {
					arg := args.Get(0).(*[]models.Todo)
					*arg = []models.Todo{
						{
							Base: models.Base{
								ID: 1,
							},
							Task: func() *string {
								s := "I am a task"
								return &s
							}(),
							Done: false,
						},
					}
				})
				f.mockQS.On("Count").Return(int64(1), nil)
			},
			args: args{
				ctx: context.Background(),
				query: models.TodoQuery{
					Todo: models.Todo{
						Task: func() *string {
							s := fake.Sentence()
							return &s
						}(),
					},
				},
			},
			want: []models.Todo{
				{
					Base: models.Base{
						ID: 1,
					},
					Task: func() *string {
						s := "I am a task"
						return &s
					}(),
					Done: false,
				},
			},
			want1:   int64(1),
			wantErr: false,
		},
		{
			name:   "with_invalid_data_should_fail",
			fields: fields{},
			prepare: func(f *fields) {
				f.mockDB.On("QueryTable", mock.Anything).Return(f.mockQS)
				f.mockQS.On("OrderBy", mock.Anything).Return(f.mockQS)
				f.mockQS.On("Filter", mock.Anything, mock.Anything).Return(f.mockQS)
				f.mockQS.On("All", mock.Anything).Return(int64(0), orm.ErrNoRows)
			},
			args: args{
				ctx: context.Background(),
				query: models.TodoQuery{
					Todo: models.Todo{
						Task: func() *string {
							s := fake.Sentence()
							return &s
						}(),
					},
				},
			},
			want: func() []models.Todo {
				var empty []models.Todo
				return empty
			}(),
			want1:   int64(0),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields = fields{
				mockDB: mocks.NewOrmerMock(t),
				mockQS: mocks.NewQuerySeterMock(t),
			}

			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}

			SUT := NewTodoRepo(tt.fields.mockDB)
			got, got1, err := SUT.FindAll(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoRepo.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoRepo.FindAll() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("todoRepo.FindAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
