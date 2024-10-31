
"""
Calendar

-months: List<Month>
-days: List<Day>
-current_year: Year
-----------------------
+Calendar(current_year)
-init_months()
-init_days()
"""

class Calendar:
    def __init__(self, current_year: int):
        self.current_year = Year(current_year)
        self.months: list[Month] = []
        self.days: list[Day] = []
        self._init_months()
        self._init_days()

    def _init_months(self):
        month_names = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
        month_days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31] # Feb defaults to 28 days
        if self.current_year.is_leap_year():
            month_days[1] = 29

        n = len(month_names)
        for i in range(n):
            self.months.append(Month(month_names[i], month_days[i]))

    def _init_days(self):
        n = len(self.months)
        for i in range(n):
            cur_month = self.months[i]
            for d in range(1, cur_month.num_days+1):
                cur_day = Day(d, i+1, self.current_year.get_year())
                self.days.append(cur_day)

    @staticmethod
    def _zellers_congruence(day: int, month: int, year: int):
        """
        Calculates the day of the week for a given date using Zeller's Congruence.
        Returns:
            0 = Saturday, 1 = Sunday, 2 = Monday, ..., 6 = Friday
        """
        if month < 3:
            month += 12
            year -= 1
        K = year % 100
        J = year // 100
        h = (day + 13*(month+1)//5 + K + K//4 + J//4 + 5*J) % 7
        return h
    
    def display_calendar(self):
        """
        Display the calendar month by month
        """
        for month in self.months:
            print(f"\n{month.name} {self.current_year.get_year()} {month.num_days} days")
            print("Su Mo Tu We Th Fr Sa")
            # determine the first weekday of the month
            first_weekday = self._zellers_congruence(1, self.months.index(month)+1, self.current_year.get_year())
            # Adjust Zeller's output to 0=Sunday, 6=Saturday
            # Zeller's: 0=Saturday, 1=Sunday, ..., 6=Friday
            # Adjust to: 0=Sunday, 1=Monday, ..., 6=Saturday
            zeller_to_calendar = {0: 6, 1: 0, 2: 1, 3: 2, 4: 3, 5: 4, 6: 5}
            start_day = zeller_to_calendar[first_weekday]

            # Init the day counter
            day_counter = 1
            # print leading spaces for the first week
            days_str = "   " * start_day

            while day_counter <= month.num_days:
                days_str += f"{day_counter:2d} "
                start_day += 1
                if start_day == 7:
                    days_str += "\n"
                    start_day = 0
                day_counter += 1
            print(days_str)



"""
Year

-year: int
-----------------
+Year(year)
-get_year(): int
-set_year(year)
-is_leap_year(): bool
"""

class Year:
    def __init__(self, year: int):
        self.year = year

    def get_year(self):
        return self.year
    
    def set_year(self, year):
        self.year = year

    def is_leap_year(self):
        return (self.year % 4 == 0 and self.year % 100 != 0) or (self.year % 400 == 0)
    
    def __repr__(self):
        return f"Year({self.year})"

"""
Month

-name: str
-num_days: int
-----------------------
+Month(name, num_days)
"""

class Month:
    def __init__(self, name: str, num_days: int):
        self.name = name
        self.num_days = num_days

    def __repr__(self):
        return f"Month({self.name}, {self.num_days} days)"

"""
Day

-day_of_month: int
-month: int
-year: int
------------------------------
+Day(day_of_month, month, year)
"""

class Day:
    def __init__(self, day_of_month: int, month: int, year: int):
        self.day_of_month = day_of_month
        self.month = month
        self.year = year

    def __repr__(self):
        return f"Day({self.day_of_month}, {self.month}, {self.year})"