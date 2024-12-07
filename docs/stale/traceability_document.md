
# Traceability Document

## Purpose

This document aims to provide a clear traceability matrix linking software requirements to their corresponding design artifacts.\
It ensures that each functional and non-functional requirement is addressed by specific software artifacts (use case diagrams, class diagrams, and activity diagrams).\
Facilitating validation and alignment throughout the development process.

---

## Use Case Diagram Traceability

| Artifact ID |          Artifact Name           |       Requirement ID       |
|:-----------:|:--------------------------------:|:--------------------------:|
|     UC1     |      Meeting Room Use Case       | FR7.1, FR7.2, FR7.4, FR7.8 |
|     UC2     | Creating a Meeting Room Use Case | FR1.1, FR1.4, FR7.6, FR7.7 |
|     UC3     |      Start Meeting Use Case      |    FR7.1, FR7.4, FR7.8     |
|     UC4     |   Manage Permissions Use Case    |        FR7.2, FR7.6        |
|     UC5     |      Join Meeting Use Case       |           FR7.4            |
|     TBD     |               TBD                |    FR7.3, FR7.5, NFR7.7    |

---

## Class Diagram Traceability

|   Artifact Name   |                  Requirement ID                   |
|:-----------------:|:-------------------------------------------------:|
| MeetingController | FR1.1, FR1.2, FR1.3, FR1.4, FR1.5, NFR1.3, NFR1.5 |
|  UserController   |    FR2.1, FR2.2, FR2.3, FR2.5, NFR2.1, NFR2.3     |
|   MeetingModel    |            FR1.1, FR1.4, FR1.5, NFR1.5            |
|      Client       |    FR3.1, FR3.2, FR3.3, FR3.5, NFR3.1, NFR3.5     |
|        Hub        |        FR6.1, FR6.2, FR6.3, FR6.5, NFR6.4         |
|       Room        |    FR4.1, FR4.2, FR4.3, FR4.4, NFR4.2, NFR4.5     |
|       Peer        |        FR5.1, FR5.2, FR5.5, NFR5.2, NFR5.4        |

---

## Activity Diagram Traceability

| Artifact ID |        Artifact Name         |       Requirement ID        |
|:-----------:|:----------------------------:|:---------------------------:|
|     AD1     |    Start Meeting Activity    |     FR7.1, FR7.4, FR7.7     |
|     AD2     |    Join Meeting Activity     |        FR7.4, FR7.8         |
|     AD3     | Create Meeting Room Activity | FR1.1, FR1.4, FR1.5, NFR1.5 |
|     AD4     | Manage Permissions Activity  |    FR7.2, FR7.6, NFR7.1     |
|     AD5     |    Leave Meeting Activity    |            FR7.3            |

---

## Links to Artifacts

Below are links to the referenced artifacts for further details:

* [Use Case Diagrams](../../artifacts/XsQuadrant%20Activity%20Diagrams.drawio.pdf)
* [Class Diagrams](../../artifacts/New%20Class%20Diagrams.pdf)
* [Activity Diagrams](../../artifacts/XsQuadrant%20Activity%20Diagrams.drawio.pdf)
