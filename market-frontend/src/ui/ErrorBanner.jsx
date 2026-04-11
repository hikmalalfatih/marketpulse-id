function ErrorBanner({ message, onRetry }) {
  return (
    <div className="error-banner">
      <span className="icon">⚠️</span>
      <span>{message}</span>
      {onRetry && (
        <button onClick={onRetry} className="retry-btn">
          Retry
        </button>
      )}
    </div>
  )
}

export { ErrorBanner }

