package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Contract struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Endpoint string `json:"endpoint"`
	Method string `json:"method"`
	ExpectedStatus int `json:"expected_status"`
	ExpectedBody string `json:"expected_body"`
	Headers string `json:"headers"`
	LastResult string `json:"last_result"`
	LastRunAt string `json:"last_run_at"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"assay2.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS contracts(id TEXT PRIMARY KEY,name TEXT NOT NULL,endpoint TEXT DEFAULT '',method TEXT DEFAULT 'GET',expected_status INTEGER DEFAULT 200,expected_body TEXT DEFAULT '',headers TEXT DEFAULT '{}',last_result TEXT DEFAULT 'pending',last_run_at TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Contract)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO contracts(id,name,endpoint,method,expected_status,expected_body,headers,last_result,last_run_at,created_at)VALUES(?,?,?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Endpoint,e.Method,e.ExpectedStatus,e.ExpectedBody,e.Headers,e.LastResult,e.LastRunAt,e.CreatedAt);return err}
func(d *DB)Get(id string)*Contract{var e Contract;if d.db.QueryRow(`SELECT id,name,endpoint,method,expected_status,expected_body,headers,last_result,last_run_at,created_at FROM contracts WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Endpoint,&e.Method,&e.ExpectedStatus,&e.ExpectedBody,&e.Headers,&e.LastResult,&e.LastRunAt,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Contract{rows,_:=d.db.Query(`SELECT id,name,endpoint,method,expected_status,expected_body,headers,last_result,last_run_at,created_at FROM contracts ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Contract;for rows.Next(){var e Contract;rows.Scan(&e.ID,&e.Name,&e.Endpoint,&e.Method,&e.ExpectedStatus,&e.ExpectedBody,&e.Headers,&e.LastResult,&e.LastRunAt,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM contracts WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM contracts`).Scan(&n);return n}
