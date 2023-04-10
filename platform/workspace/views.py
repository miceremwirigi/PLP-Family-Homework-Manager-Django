from django.shortcuts import render, redirect
import calendar
from calendar import HTMLCalendar
from datetime import datetime, timedelta
from members.views import login_user


# Create your views here.


def home(request, year=datetime.now().year, month=datetime.now().strftime('%B'), ):
    if not request.user.is_authenticated:
        return redirect (login_user)

    else:
        title = "Pending Assignments"
        assignments = {"Biochemistry": datetime.now()+ timedelta(days=7),
            "Programming":datetime.now()+ timedelta(days=7),
            "Calculus":datetime.now()+ timedelta(days=7),
            "Linear Algebra":datetime.now()+ timedelta(days=7),
            "Complex Analysis":datetime.now()+ timedelta(days=7),
        }

        print(assignments["Calculus"])

        month = month.title()

        # Convert month from name to number
        month_number = list(calendar.month_name).index(month)
        month_number = int(month_number)

        # Create a calendar
        cal = HTMLCalendar().formatmonth(year, month_number)

        # Current year
        now = datetime.now()
        current_year = now.year
        print(current_year)
        current_month = now.strftime("%B")
        current_day = now.strftime("%A")

        # Current time
        time = now.strftime("%m/%d/%Y %I:%M:%S GMT")

        context = {
            "workspace": title,
            "year": year,
            "current_day": current_day,
            "current_month": current_month,
            "current_year": current_year,
            "month_number": month_number,
            "cal": cal,        
            "time": time,
            "assignments": assignments
        }

        return render(request, 'workspace/home.html', context)
