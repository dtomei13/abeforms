import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import { Card, Container, Image, Button } from "semantic-ui-react";
import { Redirect } from "react-router-dom";
let endpoint = "http://localhost:8080";
class LawyerDashboard extends Component {
  constructor(props) {
    super(props);
    this.state = {
      case: "",
      open_items: [],
      my_items: [],
      redirect: "",
    };
  }
  componentDidMount() {
    this.getOpenCase();
    this.getMyCase();
  }
  getOpenCase = () => {
    axios.get(endpoint + "/lawyerdashboard/api/opencases").then((res) => {
      console.log(res);
      if (res.data) {
        console.log("setting open cases info");
        this.setState({
          open_items: res.data.map((item) => {
            return (
              <Card>
                <Card.Content>
                  <Image
                    floated="right"
                    size="mini"
                    src="https://react.semantic-ui.com/images/avatar/large/molly.png"
                  />
                  <Card.Header>
                    {item.firstname + " " + item.lastname}
                  </Card.Header>
                  <Card.Meta>{item.stateofissue}</Card.Meta>
                  <Card.Description>{item.description}</Card.Description>
                </Card.Content>
                <Card.Content extra>
                  <div className="ui two buttons">
                    <Button
                      basic
                      color="green"
                      onClick={() => this.caseComplete(item._id)}
                    >
                      Accept
                    </Button>
                    <Button basic color="red">
                      Decline
                    </Button>
                  </div>
                </Card.Content>
              </Card>
            );
          }),
        });
        console.log(this.setState);
      } else {
        this.setState({
          open_items: [],
        });
      }
    });
  };
  caseComplete = (id) => {
    axios
      .put(endpoint + "/lawyerdashboard/api/takecase/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      })
      .then(() => {
        this.getMyCase();
        this.getOpenCase();
      });
  };
  getMyCase = () => {
    axios.get(endpoint + "/lawyerdashboard/api/mycases").then((res) => {
      if (res.data) {
        console.log("setting my cases info");
        this.setState({
          my_items: res.data.map((item) => {
            console.log(item);
            return (
              <Card>
                <Card.Content>
                  <Image
                    floated="right"
                    size="mini"
                    src="https://react.semantic-ui.com/images/avatar/large/molly.png"
                  />
                  <Card.Header>
                    {item.firstname + " " + item.lastname}
                  </Card.Header>
                  <Card.Meta>{item.stateofissue}</Card.Meta>
                  <Card.Description>{item.description}</Card.Description>
                </Card.Content>
                <Card.Content extra>
                  <div className="ui buttons">
                    <Button basic color="grey">
                      Details
                    </Button>
                  </div>
                </Card.Content>
              </Card>
            );
          }),
        });
        console.log(this.setState);
        //console.log(this.setState)
      } else {
        this.setState({
          my_items: [],
        });
      }
    });
  };
  render() {
    try {
      var userEmail = this.props.location.state.email;
      var accessToken = this.props.location.state.access_token;
      var refreshToken = this.props.location.state.refresh_token;
      var expiry = this.props.location.state.expiry;
      console.log(expiry);
      console.log(userEmail);
      console.log("HERE");
    } catch (e) {
      return <Redirect to={"/lawyerdashboard/sign_in"} />; //Check if user is authenticated
    }

    return (
      <Container>
        <div className="App">
          <h1>Open Cases</h1>
          <div className="row">
            <Card.Group>{this.state.open_items}</Card.Group>
          </div>

          <h1>My Cases</h1>
          <div className="row">
            <Card.Group>{this.state.my_items}</Card.Group>
          </div>
        </div>
      </Container>
    );
  }
}
export default LawyerDashboard;
