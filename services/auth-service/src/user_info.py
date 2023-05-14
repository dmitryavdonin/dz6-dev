class UserInfo(object):
    def __init__(self, id, login, password, name, middleName,surname, phone, city, role, created_at, modified_at):
        self.id = id
        self.login = login
        self.password = password
        self.name = name
        self.middleName = middleName
        self.surname = surname
        self.phone = phone
        self.city = city
        self.role = role
        self.created_at = created_at
        self.modified_at = modified_at