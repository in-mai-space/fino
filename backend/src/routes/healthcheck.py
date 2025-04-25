from fastapi import APIRouter

healthcheck_router = APIRouter(
    prefix="/healthcheck",
    tags=["healthcheck"]
)

@healthcheck_router.get("/")
def health_check():
    return {"status": "healthy", "version": "1.0.0"}