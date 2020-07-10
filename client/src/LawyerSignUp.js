import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import { Container } from "semantic-ui-react";
import styles from './LawyerSignUpStyle.module.css'
import abeLogo from "./abeLogo.png";

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
    var val = true;
    var err = [];
    // check if email is in the correct format
    if(!this.state.EmailAddress.match(/^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/)){
      val = false;
      err.push(["Invalid Email Format. The correct format is 'example@example.com'"]);
    }
    // check if first name is in the correct format
    if(!this.state.FirstName.match(/^[a-zA-Z]+[a-zA-Z-']*$/)){
      val = false;
      err.push(["\nInvalid First Name Format. Can only use letters, apostraphes and hyphens"]);
    }
    // check if last name is in the correct format
    if(!this.state.LastName.match(/^[a-zA-Z]+[a-zA-Z-']*$/)){
      val = false;
      err.push(["\nInvalid Last Name Format. Can only use letters, apostraphes and hyphens"]);
    }
    // check if phone number is in the correct format
    if(!this.state.PhoneNumber.match(/^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/)){
      val = false;
      err.push(["\nInvalid Phone Number"]);
    }
    // check if state of license is in the correct format
    if(!this.state.StateOfLicense.match(/^[a-zA-Z]+[a-zA-Z-]*$/)){
      val = false;
      err.push(["\nInvalid Input for State of License"]);
    }
    // check if expertise is in the correct format
    if(!this.state.Expertise.match(/^[a-zA-Z]+[a-zA-Z-]*$/)){
      val = false;
      err.push(["\nInvalid Input for Expertise"]);
    }
    // check if password is in the correct format
    if(!this.state.Password.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{8,})/)){
      val = false;
      err.push(["\nInvalid Password. The password should have atleast 1 Upper Case, 1 Lower Case, 1 special character. It should be atleast 8 characters long"]);
    }
    // check if retype password is in the correct format
    if(!this.state.RetypePassword.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{8,})/)){
      if(this.state.RetypePassword != this.state.Password){
        val = false;
        err.push(["\nRetype Password is not the same as Password"]);
      }
    }
    if (val){
      axios
          .post(
              endpoint + "/lawyerdashboard/sign_up/api/signup",
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
    }else{
      alert(err);
    }
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
              <div className={styles.frm}>
                <h1>Create your Abe Legal Account</h1>
                <form onSubmit={this.onSubmit}>
                  <div class={styles.formGroup}>
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

                  <div class={styles.formGroup}>
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

                  <div class={styles.formGroup}>
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

                  <div class={styles.formGroup}>
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

                  <div class={styles.formGroup}>
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

                  <div class={styles.formGroup}>
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

                  <div class={styles.formGroup}>
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

                  <div class={styles.formGroup}>
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

                  <div class={styles.formGroup}>
                    <button type="submit" class="btn btn-success btn-lg">
                      Submit
                    </button>
                    <p>
                    Already have an account? Click{" "}
                    <a href="/lawyerdashboard/sign_in">here</a>
                  </p>
                  </div>

                </form>
              </div>
            </div>
          </div>
        </Container>
    );
  }
}

export default LawyerSignUp;
