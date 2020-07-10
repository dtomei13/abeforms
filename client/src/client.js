import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import abeLogo from "./abeLogo.png";
import styles from './clientsStyle.module.css';

let endpoint = "http://localhost:8080";

class Client extends Component {
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
        endpoint + "/client/api/client",
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
          <img src={abeLogo} className={styles.logo} alt="logo"></img>
          <div className="App">
            <div className="container" id="registration-form">
              <div></div>

              <div className={styles.frm}>
                <h1 className={styles.CF}>Case Form</h1>
                <form className="containerWithoutTitle" onSubmit={this.onSubmit}>
                  <div className={styles.formGroup}>
                    <h5>First Name:</h5>
                    <div>
                      <input
                          type="text"
                          className="form-control"
                          placeholder="Enter first name"
                          name= "FirstName"
                          id="firstName"
                          onChange={this.handleChange}
                          value = {FirstName || ''}

                      />
                    </div>
                  </div>

                  <p></p>

                  <div className={styles.formGroup}>
                    <h5>Last Name:</h5>
                    <div>
                      <input
                          type="text"
                          className="form-control"
                          placeholder="Enter last name"
                          name="LastName"
                          id="lastName"
                          onChange={this.handleChange}
                          value = {LastName || ''}

                      />
                    </div>
                  </div>


                  <p></p>

                  <div className={styles.formGroup}>
                    <h5>Email:</h5>
                    <div>
                      <input
                          type="text"
                          className="form-control"
                          placeholder="Enter email"
                          name="EmailAddress"
                          id="emailAddress"
                          onChange={this.handleChange}
                          value = {EmailAddress || ''}

                      />
                    </div>
                  </div>

                  <p></p>
                  <div className={styles.formGroup}>
                    <h5>Phone Number:</h5>
                    <div>
                      <input
                          type="text"
                          className="form-control"
                          placeholder="Enter phone number"
                          name="PhoneNumber"
                          id="phoneNumber"
                          onChange={this.handleChange}
                          value = {PhoneNumber || ''}

                      />
                    </div>
                  </div>

                  <p></p>
                  <div className={styles.formGroup}>
                    <h5>Location of Legal Issue:</h5>
                    <div>
                      <input
                          type="text"
                          className="form-control"
                          placeholder="Enter location of legal issue"
                          name="StateOfIssue"
                          id="locationOfLegalIssue"
                          onChange={this.handleChange}
                          value = {StateOfIssue || ''}

                      />
                    </div>
                  </div>
                  <p></p>
                  <div className={styles.formGroup}>
                    <h5>Description:</h5>
                    <div>
                      <input
                          type="text"
                          className={styles.formStyle}
                          placeholder="Enter description of legal issue"
                          name="Description"
                          id="Description"
                          onChange={this.handleChange}
                          value = {Description || ''}

                      />
                    </div>
                  </div>
                  <p></p>
                  <div className={styles.formGroup}>
                    <h5>How did you find Abe Legal?</h5>
                    <div>
                      <input
                          type="text"
                          className="form-control"
                          placeholder="Enter the way you find Abe Legal"
                          name="FindHow"
                          id="FindHow"
                          onChange={this.handleChange}
                          value = {FindHow || ''}

                      />
                    </div>
                  </div>
                  <p></p>
                  <div className={styles.formGroup}>
                    <h5>Share your Social Media: (Instagram, Twitter, etc)</h5>
                    <div>
                      <input
                          type="text"
                          className="form-control"
                          placeholder="Enter your social media ID "
                          name="SocialMedia"
                          id="SocialMedia"
                          onChange={this.handleChange}
                          value = {SocialMedia || ''}

                      />
                    </div>
                  </div>
                  <p></p>
                  <button type="submit" className={styles.submit}>
                    Submit
                  </button>
                </form>
              </div>
            </div>
          </div>
        </div>
    );
  }
}

export default Client;
