import json
import aiohttp_jinja2
from aiohttp import web, hdrs
import requests

from src.user_info import UserInfo

routes = web.RouteTableDef()


@routes.route(hdrs.METH_GET, '/sign-out')
async def login(req: web.Request):
    storage = req.app.storage
    session_id = req.cookies.get('session_id')
    storage.remove_session(session_id)
    return web.Response(text='Logged out')


@routes.route(hdrs.METH_GET, '/sign-in')
async def login(req: web.Request):
    state_key = req.query.get('state')
    return aiohttp_jinja2.render_template('login.html', req, dict(state_key=state_key))


@routes.route(hdrs.METH_POST, '/sign-in')
async def login(req: web.Request):
    storage = req.app.storage
    data = await req.json()
    login = data.get('login')
    pwd = data.get('pass')
    state_key = data.get('state')  
    state = storage.pop_state(state_key)            
    if state:
        response = web.HTTPFound(state['req_url'])
    else:            
        if login and pwd:                  
            credentials = {'login': login, 'pass': pwd}
            baseUrl = req.app.config.getUserServiceUri()                
            resp = requests.post(baseUrl + "/user/creds", json=credentials)    
            if resp:
                ui = UserInfo(**resp.json())
                if ui:
                    session_id = storage.create_session(username=login, user_id=ui.id)
                    jsonSessionId = {'session_id':session_id}
                    response = web.Response(status=200, text= json.dumps(jsonSessionId))
                    response.set_cookie('session_id', session_id)
                else:
                    response = web.HTTPNotFound(text="User " + login + " not found")    
            else:
                response = web.HTTPInternalServerError(text='No response from profile service')
        else:
            response = web.Response(status=400, text='No username or password specified')                    
    return response
