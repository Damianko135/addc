// Unless specified, it will think it is dev... Should specify it somewhere in the build process
const environment = (window?.ENV || 'development'); // safer if injected at runtime

  const link = document.createElement('link');
  link.rel = 'stylesheet';
  link.href = environment === 'development' 
      ? 'style.css?v=' + new Date().getTime() // Always fresh
      : 'style.css'; // Cached in production
  document.head.appendChild(link);
