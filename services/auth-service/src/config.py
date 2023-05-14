from os import environ as env

def setup_config(app):
    app.config = Config()
    
class Config:
    
    def __init__(self):
        pass
    
    def getAppPort(self):
        if env['APP_PORT']:
            return 8080 #env['APP_PORT']
        else:
            return 8080
        
    def getUserServiceUri(self):
        if env['USERS_SERVICE_URI']:
            return env['USERS_SERVICE_URI']
        else:
            return ""