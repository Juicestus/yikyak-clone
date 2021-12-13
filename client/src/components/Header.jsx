import logo from '../logo.svg';
import PostButton from './PostButton.jsx';

function Header(props) {
    return (
        <header className="app-header">
          <img src={logo} className="logo" alt="logo" /> 
          <h1>{props.title}</h1>
          <PostButton />
        </header>
    );
  }
  
  export default Header;