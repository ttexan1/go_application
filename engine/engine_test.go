package engine

type testStorage struct{}

func (f *testStorage) Close()                        {}
func (f *testStorage) DropTables()                   {}
func (f *testStorage) Migrate()                      {}
func (f *testStorage) NewArticleRepo() ArticleRepo   { return nil }
func (f *testStorage) NewCategoryRepo() CategoryRepo { return nil }
func (f *testStorage) NewWriterRepo() WriterRepo     { return nil }
