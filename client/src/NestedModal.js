import React, { Component } from "react";
import "./App.css";
import { Button, Modal, Form, Icon, Header } from "semantic-ui-react";
import DayPicker from "react-day-picker";
import "antd/dist/antd.css";
import { TimePicker } from "antd";

// Or import the input component
import DayPickerInput from "react-day-picker/DayPickerInput";

import "react-day-picker/lib/style.css";

class NestedModal extends Component {
  constructor(props) {
    super(props);
    this.handleDayClick = this.handleDayClick.bind(this);
    this.state = {
      selectedDay: undefined,
    };
  }

  state = { open: false };

  open = () => this.setState({ open: true });
  close = () => this.setState({ open: false });

  handleDayClick(day, { selected }) {
    if (selected) {
      // Unselect the day if already selected
      this.setState({ selectedDay: undefined });
      return;
    }
    this.setState({ selectedDay: day });
  }

  render() {
    const { open } = this.state;

    return (
      <Modal
        open={open}
        onOpen={this.open}
        onClose={this.close}
        size="small"
        trigger={
          <Button primary icon>
            Schedule a Zoom meeting <Icon name="right chevron" />
          </Button>
        }
      >
        <Modal.Header>Schedule a Zoom Meeting with your Client</Modal.Header>
        <Modal.Content>
          <Form>
            <Form.Field>
              <label>Preferred Date</label>
              <DayPicker
                className="preferredDate"
                disabledDays={{ daysOfWeek: [0, 6] }}
                onDayClick={this.handleDayClick}
                selectedDays={this.state.selectedDay}
              />
              {this.state.selectedDay ? (
                <p>
                  You selected {this.state.selectedDay.toLocaleDateString()}
                </p>
              ) : (
                <p>Please select a day.</p>
              )}
              {/*<input
                type="date"
                className="preferredDate"
                placeholder="Preferred Date"
              />*/}
            </Form.Field>
            <Form.Field>
              <label>Preferred Time (First Choice)</label>
              <TimePicker
                className="firstTime"
                use12Hours
                minuteStep={15}
                format="h:mm a"
                style={{ width: 140 }}
              />
              {/*  <select type="checkbox" className="timesFirst">
                <option value="8">8:00 AM</option>
                <option value="9">9:00 AM</option>
                <option value="10">10:00 AM</option>
                <option value="11">11:00 AM</option>
                <option value="12">12:00 Noon</option>
                <option value="1">1:00 PM</option>
                <option value="2">2:00 PM</option>
                <option value="3">3:00 PM</option>
              </select>*/}
            </Form.Field>
            <Form.Field>
              <label>Preferred Time (Second Choice)</label>
              <TimePicker
                className="firstTime"
                use12Hours
                minuteStep={15}
                format="h:mm a"
                style={{ width: 140 }}
              />
              {/*<select type="checkbox" className="timesSecond">
                <option value="8">8:00 AM</option>
                <option value="9">9:00 AM</option>
                <option value="10">10:00 AM</option>
                <option value="11">11:00 AM</option>
                <option value="12">12:00 Noon</option>
                <option value="1">1:00 PM</option>
                <option value="2">2:00 PM</option>
                <option value="3">3:00 PM</option>
              </select> */}
            </Form.Field>
            <Form.Field>
              <label>Preferred Time (Third Choice)</label>
              <TimePicker
                className="firstTime"
                use12Hours
                minuteStep={15}
                format="h:mm a"
                style={{ width: 140 }}
              />
              {/*       <select type="checkbox" className="timesThird">
                <option value="8">8:00 AM</option>
                <option value="9">9:00 AM</option>
                <option value="10">10:00 AM</option>
                <option value="11">11:00 AM</option>
                <option value="12">12:00 Noon</option>
                <option value="1">1:00 PM</option>
                <option value="2">2:00 PM</option>
                <option value="3">3:00 PM</option>
              </select> */}
            </Form.Field>
            <Form.Field>
              <Header as="h4">What to Expect for your Consultation</Header>
              <ul>
                <li>
                  You will be able to speak with your client through Zoom for
                  30-60 minutes
                </li>

                <li>
                  There will be a 5-minute period at the end of the consultation
                  for your client to ask questions
                </li>
                <li>
                  The consultation will be recorded for your client's future
                  reference
                </li>
              </ul>
            </Form.Field>

            <Form.Field>
              <Button
                floated="right"
                color="blue"
                type="submit"
                icon="check"
                content="Submit"
                onClick={this.close}
              />
            </Form.Field>
          </Form>
        </Modal.Content>
      </Modal>
    );
  }
}

export default NestedModal;