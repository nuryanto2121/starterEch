package queryfileupload

const (
	QuerySave = `
	INSERT INTO public.sa_file_upload (file_name, file_path, file_type, created_by, updated_by) 
	VALUES(:file_name, :file_path, :file_type, :created_by, :updated_by);

	`
)
