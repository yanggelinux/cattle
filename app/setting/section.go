package setting

type APPSettingS struct {
	Name        string
	Service     string
	LogPath     string
	LogLevel    string
	Development bool
}

type ServerSettingS struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  int
	WriteTimeout int
}

type JWTSettingS struct {
	JwtSecret string
	JwtExpire int
	JwtIssuer string
}

type MySQLSettingS struct {
	Host            string
	Username        string
	Password        string
	Port            int
	DBName          string
	SecondDBName    string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

type CronSettingS struct {
	Status int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
