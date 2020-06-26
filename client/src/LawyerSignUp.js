import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import { Container } from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class LawyerSignUp extends Component {
  constructor(props) {
    super(props);
    this.state = {
      FirstName: "",
      LastName: "",
      PhoneNumber: "",
      EmailAddress: "",
      StateOfLicense: "",
      Expertise: "",
      Password: "",
      RetypePassword: "",
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
      StateOfLicense,
      Expertise,
      Password,
      RetypePassword,
    } = this.state;
    axios
      .post(
        endpoint + "/lawyerdashboard/api/signup",
        {
          FirstName: FirstName,
          LastName: LastName,
          PhoneNumber: PhoneNumber,
          EmailAddress: EmailAddress,
          StateOfLicense: StateOfLicense,
          Expertise: Expertise,
          Password: Password,
          RetypePassword: RetypePassword,
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
      StateOfLicense,
      Expertise,
      Password,
      RetypePassword,
    } = this.state;
    return (
      <Container>
        <div className="App">
          <div className="container" id="registration-form">
            <div className="image"></div>
            <div className="frm">
              <h1>Create your Abe Legal Account</h1>
              <form onSubmit={this.onSubmit}>
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
                  <h5>State Of License: </h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter location of legal issue"
                      name="StateOfLicense"
                      id="StateOfLicense"
                      onChange={this.handleChange}
                      value={StateOfLicense || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Expertise:</h5>
                  <div>
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Enter expertise of legal area"
                      name="Expertise"
                      id="Expertise"
                      onChange={this.handleChange}
                      value={Expertise || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Password: (Minimum 8 Characters)</h5>
                  <div>
                    <input
                      type="password"
                      class="form-control"
                      name="Password"
                      id="Password"
                      onChange={this.handleChange}
                      value={Password || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <h5>Retype Password: </h5>
                  <div>
                    <input
                      type="password"
                      class="form-control"
                      name="RetypePassword"
                      id="RetypePassword"
                      onChange={this.handleChange}
                      value={RetypePassword || ""}
                    />
                  </div>
                </div>

                <div class="form-group">
                  <button type="submit" class="btn btn-success btn-lg">
                    Submit
                  </button>
                </div>
                <p>
                  Already have an account? Click{" "}
                  <a href="/lawyerdashboard/sign_in">here</a>
                </p>
              </form>
            </div>
          </div>
        </div>
      </Container>
    );
  }
}

export default LawyerSignUp;
