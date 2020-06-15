import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import abeLogo from "./abeLogo.png";

let endpoint = "http://localhost:8080";

class Clients extends Component {
  constructor(props) {
    super(props);
    this.state = {
      FirstName: "",
      LastName: "",
      PhoneNumber: "",
      EmailAddress: "",
      Description: "",
      StateOfIssue: "",
    };
  }
  handleChange = (event) => {
    this.setState({
      [event.target.name]: event.target.value,
    });
  };

  onSubmit = () => {
    const {
      FirstName,
      LastName,
      PhoneNumber,
      EmailAddress,
      Description,
      StateOfIssue,
      FindHow,
      SocialMedia,
    } = this.state;
    axios
      .post(
        endpoint + "/api/client",
        {
          FirstName: FirstName,
          LastName: LastName,
          PhoneNumber: PhoneNumber,
          EmailAddress: EmailAddress,
          Description: Description,
          StateOfIssue: StateOfIssue,
          findHow: FindHow,
          socialMedia: SocialMedia,
        },
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      )
      .then((res) => console.log(FirstName));
  };

  render() {
    const {
      FirstName,
      LastName,
      PhoneNumber,
      EmailAddress,
      Description,
      StateOfIssue,
      FindHow,
      SocialMedia,
    } = this.state;
    return (
      <div>
        <img src={abeLogo} className="logo" alt="logo"></img>
        <div className="App">
          <div className="container" id="registration-form">
            <div className="image"></div>

            <div className="frm">
              <h1>Case Form</h1>
              <form className="containerWithoutTitle" onSubmit={this.onSubmit}>
                <div class="form-group">
                  <h5>First Name:</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter first name"
                      name="FirstName"
                      id="firstName"
                      onChange={this.handleChange}
                      value={FirstName || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Last Name:</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter last name"
                      name="LastName"
                      id="lastName"
                      onChange={this.handleChange}
                      value={LastName || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Email:</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter email"
                      name="EmailAddress"
                      id="emailAddress"
                      onChange={this.handleChange}
                      value={EmailAddress || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Phone Number:</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter phone number"
                      name="PhoneNumber"
                      id="phoneNumber"
                      onChange={this.handleChange}
                      value={PhoneNumber || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Location of Legal Issue:</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter location of legal issue"
                      name="StateOfIssue"
                      id="locationOfLegalIssue"
                      onChange={this.handleChange}
                      value={StateOfIssue || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Description:</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter description of legal issue"
                      name="Description"
                      id="Description"
                      onChange={this.handleChange}
                      value={Description || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>How did you find Abe Legal?</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter the way you find Abe Legal"
                      name="FindHow"
                      id="FindHow"
                      onChange={this.handleChange}
                      value={FindHow || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Share your Social Media: (Instagram, Twitter, etc)</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter your social media ID "
                      name="SocialMedia"
                      id="SocialMedia"
                      onChange={this.handleChange}
                      value={SocialMedia || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <button type="submit" class="btn btn-success btn-lg">
                    Submit
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default Clients;
