package postgres

import (
	"context"

	"github.com/admin5225/coursework-DB.-managing-campaigns-of-the-city/internal/application/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, application *domain.Application) error {
	query := `
		CALL main.add_application($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		application.Description,
		application.HouseID,
		application.SpecialistID,
		application.WorkTypeID,
		application.StatusID,
	)

	return err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	query := `
		CALL main.delete_application($1)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}

func (r *Repository) Close(ctx context.Context, id int) error {
	query := `
		CALL main.close_application($1)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}

func (r *Repository) GetCloseApplications(ctx context.Context) ([]domain.FullApplication, error) {
	query := `
		SELECT * FROM main.get_closed_applications()
	`
	rows, err := r.db.Query(
		ctx,
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	applications := make([]domain.FullApplication, 0)

	for rows.Next() {

		var application domain.FullApplication

		err := rows.Scan(
			&application.ID,
			&application.Description,
			&application.DateTime,
			&application.Street,
			&application.HouseNumber,
			&application.SpecialistName,
			&application.SpecialistPost,
			&application.WorkType,
			&application.ApplicationStatus,
		)

		if err != nil {
			return nil, err
		}

		applications = append(
			applications,
			application,
		)
	}

	return applications, nil
}

func (r *Repository) GetApplicationsStats(
	ctx context.Context,
) (*domain.Statistics, error) {

	query := `
		SELECT * FROM main.get_applications_statistics()
	`

	var stats domain.Statistics

	err := r.db.QueryRow(
		ctx,
		query,
	).Scan(
		&stats.TotalRequests,
		&stats.OpenRequests,
		&stats.ClosedRequests,
	)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *Repository) GetHouseApplications(
	ctx context.Context,
	houseID int,
) ([]domain.FullApplication, error) {

	query := `
		SELECT * FROM main.get_applications_by_house($1)
	`

	rows, err := r.db.Query(
		ctx,
		query,
		houseID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var applications []domain.FullApplication

	for rows.Next() {

		var application domain.FullApplication

		err := rows.Scan(
			&application.ID,
			&application.Description,
			&application.DateTime,
			&application.Street,
			&application.HouseNumber,
			&application.SpecialistName,
			&application.SpecialistPost,
			&application.WorkType,
			&application.ApplicationStatus,
		)

		if err != nil {
			return nil, err
		}

		applications = append(
			applications,
			application,
		)
	}

	return applications, nil
}

func (r *Repository) GetSpecialistApplications(
	ctx context.Context,
	specialistID int,
) ([]domain.FullApplication, error) {

	query := `
		SELECT * FROM main.get_applications_by_specialist($1)
	`

	rows, err := r.db.Query(
		ctx,
		query,
		specialistID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var applications []domain.FullApplication

	for rows.Next() {

		var application domain.FullApplication

		err := rows.Scan(
			&application.ID,
			&application.Description,
			&application.DateTime,
			&application.Street,
			&application.HouseNumber,
			&application.SpecialistName,
			&application.SpecialistPost,
			&application.WorkType,
			&application.ApplicationStatus,
		)

		if err != nil {
			return nil, err
		}

		applications = append(
			applications,
			application,
		)
	}

	return applications, nil
}
