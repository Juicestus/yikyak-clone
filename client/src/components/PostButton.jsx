import React, { Component } from "react";

import PostPopup from './PostPopup.jsx';

class PostButton extends Component {

  handleSubmit = (e) => {
    e.preventDefault();
    console.log('You clicked submit.');
  }

  state = { seen: false };

  togglePop = () => {
    this.setState({
      seen: !this.state.seen
    });
  };

  /*return (
    <form onSubmit={handleSubmit}>
      <button className="solid-button" type="submit">
        New Post
      </button>
    </form>
  );*/

  render() { return (
    <div>
      <a className="cool-button" onClick={this.togglePop} href="#">New Post</a>
      {this.state.seen ? <PostPopup toggle={this.togglePop} /> : null}
    </div>
  );}
}
  
export default PostButton;