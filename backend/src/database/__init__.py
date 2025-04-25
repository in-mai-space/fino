from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from src.models.user import Base

Base = declarative_base()

def connect_db(config):
    """
    Establish connection to database using SQLAlchemy
    """
    db_config = config
    
    # Construct database URL
    db_url = f"postgresql://{db_config['FINO_DB_USER']}:{db_config['FINO_DB_PASSWORD']}@{db_config['FINO_DB_HOST']}:{db_config['FINO_DB_PORT']}/{db_config['FINO_DB_NAME']}"
    
    # Connection arguments including SSL
    connect_args = {}
    if db_config['FINO_DB_SSL']:
        connect_args["sslmode"] = db_config['FINO_DB_SSL']
    
    # Create engine
    engine = create_engine(
        db_url,
        connect_args=connect_args
    )
    
    # Create session factory
    SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
    
    return {
        "engine": engine,
        "SessionLocal": SessionLocal,
        "Base": Base,
    }


def get_db(db):
    """
    Get the database object to do database operation on
    """
    session = db["SessionLocal"]()
    try:
        yield session
    finally:
        session.close()