from datetime import datetime

def current_year(request):
    current_year = datetime.now().year
    return {"current_year":current_year}
