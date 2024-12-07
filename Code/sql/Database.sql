# If error, execute twice

CREATE DATABASE College_Internal_System;
USE College_Internal_System;

#  Create the Students Table
CREATE TABLE Students (
    StudentID INT PRIMARY KEY AUTO_INCREMENT,
    FirstName VARCHAR(50),
    LastName VARCHAR(50),
    Email VARCHAR(100) UNIQUE,
    Major VARCHAR(100),
    ClassYear ENUM('Freshman', 'Sophomore', 'Junior', 'Senior', 'Plus'),
    DateOfBirth DATE,
    Gender ENUM('Male', 'Female', 'Other'),
    PhoneNumber VARCHAR(15),
    Address VARCHAR(255),
    CalculatedGrade DECIMAL(3, 2),
    CumulativeGPA DECIMAL(3, 2),
    SemesterGPA DECIMAL(3, 2)
);
 
#  Create the Courses Table
CREATE TABLE Courses (
    CourseID INT PRIMARY KEY AUTO_INCREMENT,
    CourseName VARCHAR(100),
    CourseCode VARCHAR(10) UNIQUE,
    Description TEXT,
    Credits INT,
    Department VARCHAR(100),
    InstructorID INT, -- Optional Foreign Key
    SemesterOffered VARCHAR(20),
    MaxEnrollment INT,
    FOREIGN KEY (InstructorID) REFERENCES Instructors(InstructorID)
);
 
#  Create the Prerequisites Table
CREATE TABLE Prerequisites (
    PrerequisiteID INT PRIMARY KEY AUTO_INCREMENT,
    CourseID INT,
    PrerequisiteCourseID INT,
    FOREIGN KEY (CourseID) REFERENCES Courses(CourseID),
    FOREIGN KEY (PrerequisiteCourseID) REFERENCES Courses(CourseID)
);
 
#  Create the Enrollments Table
CREATE TABLE Enrollments (
    EnrollmentID INT PRIMARY KEY AUTO_INCREMENT,
    StudentID INT,
    CourseID INT,
    Semester VARCHAR(20),
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID),
    FOREIGN KEY (CourseID) REFERENCES Courses(CourseID)
);
 
#  Create the Instructors Table (Optional)
CREATE TABLE Instructors (
    InstructorID INT PRIMARY KEY AUTO_INCREMENT,
    FirstName VARCHAR(50),
    LastName VARCHAR(50),
    Email VARCHAR(100) UNIQUE,
    Department VARCHAR(100)
);

