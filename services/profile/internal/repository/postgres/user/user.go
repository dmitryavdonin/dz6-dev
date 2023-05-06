package repository

import (
	"context"
	"errors"
	"profile/internal/domain/user"
	"profile/internal/repository/postgres/user/dao"
	"profile/pkg/tools/transaction"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = "public.user"
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"user_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateUser(ctx context.Context, user *user.User) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Insert(tableName).Columns(dao.UserColumns...).Values(user.Id(), user.Login(), user.Password(), user.Name(), user.MiddleName(), user.Surname(), user.Phone(), user.City(), user.Role(), user.CreatedAt(), user.ModifiedAt())
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateUser(ctx context.Context, id uuid.UUID, upFunc func(oldUser *user.User) (*user.User, error)) (user *user.User, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	oldUser, err := r.oneUserTx(ctx, tx, id)
	if err != nil {
		return
	}

	newUser, err := upFunc(oldUser)
	if err != nil {
		return
	}

	rawQuery := r.Builder.Update(tableName).Set("login", newUser.Login()).Set("password", newUser.Password()).Set("name", newUser.Name()).Set("middle_name", newUser.MiddleName()).Set("surname", newUser.Surname()).Set("phone", newUser.Phone()).Set("city", newUser.City()).Set("role", newUser.Role()).Set("modified_at", newUser.ModifiedAt()).Where("id = ?", newUser.Id())
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return nil, ErrUpdate
	}

	return newUser, nil
}

func (r *Repository) DeleteUserById(ctx context.Context, id uuid.UUID) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Delete(tableName).Where("id = ?", id)
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}
	return
}

func (r *Repository) ReadUserById(ctx context.Context, id uuid.UUID) (user *user.User, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	return r.oneUserTx(ctx, tx, id)
}

func (r *Repository) ReadUserByLogin(ctx context.Context, login string) (user *user.User, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Select(dao.UserColumns...).From(tableName).Where("login = ?", login)
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	daoUser, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.User])
	if err != nil {
		return nil, err
	}

	return r.toDomainUser(&daoUser)
}

func (r *Repository) oneUserTx(ctx context.Context, tx pgx.Tx, id uuid.UUID) (user *user.User, err error) {
	rawQuery := r.Builder.Select(dao.UserColumns...).From(tableName).Where("id = ?", id)
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	daoUser, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.User])
	if err != nil {
		return nil, err
	}

	return r.toDomainUser(&daoUser)
}
