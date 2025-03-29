# NextCloud Alt
## Backend Stack (proof of concept version)
- **Blob Storage**: MinIO 
- **MetaData Store**: PostgresSQL or MariaDb
- **API**: REST + WS
- **Auth**: JWT 
- **Sync Protocols**: not sure, maybe custom API or some hacky rsync setup
- **Data Integrity**: SHA256 stored in DB + MinIO ETags as secondary
- **Caching**: Redis
- **Language**: Go

## Frontend
- Anything with Crossplatform support (desktop + mobile)
