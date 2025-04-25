from sqlalchemy import Column, Integer, String, DateTime
from sqlalchemy.sql import func
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

class User(Base):
    """
    User model representing app users with their associated tokens
    
    Attributes:
    -----------
    id : int
        Primary key, unique identifier for the user
    plaid_access_token : str (encrypted)
        Token for authenticating with Plaid API
    item_id : str (encrypted)
        Plaid item identifier associated with the user's financial institution
    notion_access_token : str (encrypted)
        Token for authenticating with Notion API
    notion_refresh_token : str (encrypted)
        Token used to refresh the Notion access token when it expires
    created_at : datetime
        Timestamp when the user record was created
    updated_at : datetime
        Timestamp when the user record was last updated
    """
    __tablename__ = "users"
    
    id = Column(Integer, primary_key=True, index=True)
    plaid_access_token = Column(String, nullable=True)
    item_id = Column(String, nullable=True)
    notion_access_token = Column(String, nullable=True)
    notion_refresh_token = Column(String, nullable=True)
    created_at = Column(DateTime(timezone=True), server_default=func.now())
    updated_at = Column(DateTime(timezone=True), onupdate=func.now())
    
    def __repr__(self):
        return f"<User(id={self.id})>"
    
