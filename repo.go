package repository

import (
    "context"

    "app-service/internal/models"
    "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
    DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
    return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, user models.User) error {
    _, err := r.DB.Exec(ctx, `INSERT INTO users (name, email) VALUES ($1, $2)`, user.Name, user.Email)
    return err
}

func (r *UserRepo) GetAll(ctx context.Context) ([]models.User, error) {
    rows, err := r.DB.Query(ctx, `SELECT id, name, email FROM users`)
    if err != nil {
        return nil, err
    }

    var users []models.User
    for rows.Next() {
        var u models.User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }

    return users, nil
}

func (r *UserRepo) Update(ctx context.Context, user models.User) error {
    _, err := r.DB.Exec(ctx, `UPDATE users SET name=$1, email=$2 WHERE id=$3`,
        user.Name, user.Email, user.ID)
    return err
}

func (r *UserRepo) Delete(ctx context.Context, id int64) error {
    _, err := r.DB.Exec(ctx, `DELETE FROM users WHERE id=$1`, id)
    return err
}
