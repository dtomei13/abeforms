import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import { Container, Form } from "semantic-ui-react";
import { Redirect } from "react-router-dom";

let endpoint = "http://localhost:8080";

class LawyerSignIn extends Component {
  constructor(props) {
    super(props);
    this.state = {
      EmailAddress: "",
      Password: "",
      redirect: "",
    };
  }
  handleChange = (event) => {
    this.setState({
      [event.target.name]: event.target.value,
    });
  };

  onPress = () => {
    const { EmailAddress, Password } = this.state;
    axios
        .post(
            endpoint + "/lawyerdashboard/api/signin",
            {
              EmailAddress: EmailAddress,
              Password: Password,
            },
            {
              headers: {
                "Content-Type": "application/x-www-form-urlencoded",
              },
            }
        )
        .then((res) => {
          this.setState({
            redirect: true,
          })
        });
  };

  render() {
    const { EmailAddress, Password } = this.state;
    if (this.state.redirect) {
      return (
          <Redirect
              to={{
                pathname: "/lawyerdashboard",
              }}
          />
      );
    }
    return (
        <Container>
          <div className="App">
            <div className="container" id="registration-form">
              <div className="image"></div>
              <div className="frm">
                <h1>Sign in</h1>
                <p>to continue to Abe</p>
                <Form onSubmit={this.onPress}>
                  <div class="form-group">
                    <h5>Email Address:</h5>
                    <div>
                      <input
                          type="text"
                          class="form-control"
                          placeholder="Enter email address"
                          name="EmailAddress"
                          id="EmailAddress"
                          onChange={this.handleChange}
                          value={EmailAddress || ""}
                      />
                    </div>
                  </div>

                  <div class="form-group">
                    <h5>Password:</h5>
                    <div>
                      <input
                          type="password"
                          class="form-control"
                          placeholder="Enter password"
                          name="Password"
                          id="Password"
                          onChange={this.handleChange}
                          value={Password || ""}
                      />
                    </div>
                  </div>

                  <div class="form-group">
                    <button class="btn btn-success btn-lg">Submit</button>
                  </div>
                  <p>
                    Don't have an account? Click{" "}
                    <a href="/lawyerdashboard/sign_up">here</a>
                  </p>
                </Form>
              </div>
            </div>
          </div>
        </Container>
    );
  }
}

export default LawyerSignIn;
