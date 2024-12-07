USE College_Internal_System;
INSERT INTO Instructors VALUES (1, 'Professor', 'LastName', 'plast@wooster.edu', 'Mathematical and Computational Sciences');
INSERT INTO Instructors VALUES (2, 'Another', 'Prof', 'aprof@wooster.edu', 'Mathematical and Computational Sciences');


INSERT INTO Courses (CourseID, CourseName, CourseCode, Description, Credits, Department, InstructorID, SemesterOffered, MaxEnrollment)
VALUES (1, 'Intro to Computer Science', 101, 'An introductory Computer Science Course', 1, 'Mathematical and Computational Sciences',
1, 'Spring 2024', 30);

INSERT INTO Courses (CourseID, CourseName, CourseCode, Description, Credits, Department, InstructorID, SemesterOffered, MaxEnrollment)
VALUES (2, 'Data Structures', 200, 'A class on data structures', 1, 'Mathematical and Computational Sciences',
1, 'Spring 2024', 25);

INSERT INTO Prerequisites VALUES (1, 2, 1);

INSERT INTO Students VALUES (1, 'AStudent', 'Name', 'aname@wooster.edu', 'Computer Science', 'Junior', '2001-01-01', 'Male', '5554381929',
'3074 Address Ave.', 3.0, 3.74, 3.8);

INSERT INTO Students VALUES (2, 'Another', 'Last', 'alast@wooster.edu', 'Data Science', 'Sophomore', '2001-01-01', 'Female', '5554384929',
'3104 Address Ave.', 3.2, 3.44, 3.9);

INSERT INTO Students VALUES (3, 'Will', 'Sieber', 'wsieber@wooster.edu', 'Computer Science', 'Senior', '2003-01-01', 'Male', '3749369343',
'2308 Address Lane.', 2.0, 3.00, 1.9);

INSERT INTO Enrollments VALUES (1, 1, 1, 'Spring 2024');



SELECT * FROM Courses;
SELECT * FROM Enrollments;
SELECT * FROM Instructors;
SELECT * FROM Prerequisites;
SELECT * FROM Students;