import './styles/BigURL.scss';

const redirectHost = process.env.REACT_APP_REDIRECT_HOST;

function BigURL(props) {
  const fullRedirectURL = `${redirectHost}/${props.url}`;

  return (
    <div className="mini-url container">
      <div className="link-wrapper d-flex flex-column justify-content-center
        align-items-center mt-5">
        <h2>Hooray! Here is your bigURL.</h2>
        <img src={`${process.env.PUBLIC_URL}/images/done.gif`} alt="ready" />
        <a href={fullRedirectURL}>{fullRedirectURL}</a>
      </div>
    </div>
  );
}

export default BigURL;
