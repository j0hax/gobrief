# gobrief
Summarize iCal Events on the command line

The purpose of this program is to provide a *simple* and *fast* way to peek upcoming events in your calendar. Think of it as an `ls` command for your agenda.

## Example
```console
johannes@kirby:~ > gobrief -d 7 -u 'https://studip.uni-hannover.de/dispatch.php/ical/index/********'
Mon  2 May 10:15 Logik und formale Systeme
Mon  2 May 14:15 Mathematik 2: Analysis
Mon  2 May 14:30 Programmiersprachen und Übersetzer: Syntax and Parsing
Tue  3 May 08:15 Tutorium: Komplexität von Algorithmen
Tue  3 May 08:15 Diskrete Strukturen für Studierende der Informatik
Tue  3 May 17:30 Mathematik 2: Analysis
Wed  4 May 08:30 Proseminar E-Learning: B. Wissenschaftliches Arbeiten, 10-Cybersecurity-Paßlack
Wed  4 May 12:00 Komplexität von Algorithmen
Wed  4 May 14:00 Hardware-Praktikum
Thu  5 May 14:45 Programmierpraktikum Technische Informatik
Fri  6 May 08:00 Übung zu Mathematik 2: Analysis: 
Fri  6 May 10:15 Gruppe 7, Übung zu Diskrete Strukturen für Studierende der Informatik
Thu  5 May 19:00 Miniprojekt Besprechung
```
