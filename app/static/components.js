class NavigationBar extends HTMLElement {
    constructor() {
      super();
      const shadow = this.attachShadow({ mode: 'open' });
      shadow.innerHTML = `
        <style>
          /* Add any styles for the shadow DOM here */
          nav {
            background-color: #00227C;
            overflow: hidden;
            height: 80px;
            display: flex;
            justify-content: flex-end;
            align-items: center;
            padding: 0 20px;
          }
          nav a {
            color: #fff;
            text-decoration: none;
            padding: 0 20px;
          }
          .active {
            text-decoration: underline;
          }
        </style>
        <nav>
          <a id="home" href="/">Home</a>
          <a id="create" href="/create">Create</a>
          <a id="account" href="/account">Account</a>
        </nav>
      `;
    }
  
    connectedCallback() {
      const activePageId = this.getAttribute('page') || 'home';
      const activePageElement = this.shadowRoot.getElementById(activePageId);
  
      if (activePageElement) {
        activePageElement.classList.add('active');
      }
    }
  }
  
  customElements.define('navigation-bar', NavigationBar);
  