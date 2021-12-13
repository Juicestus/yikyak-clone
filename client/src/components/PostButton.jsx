import React, { Component } from "react";

import PostPopup from "./PostPopup.jsx";

class PostButton extends Component {
  handleSubmit = (e) => {
    e.preventDefault();
    console.log("You clicked submit.");
  };

  state = { seen: false };

  togglePop = () => {
    this.setState({
      seen: !this.state.seen,
    });
  };

  render() {
    return (
      <div>
        <a className="transparent-button open-new-post-window" onClick={this.togglePop}>
          {!this.state.seen ? "New Post" : "Cancel"}
        </a>
        {this.state.seen ? <PostPopup toggle={this.togglePop} /> : null}
      </div>
    );
  }
}

export default PostButton;
