function Skeleton({ rows = 5, type = 'table' }) {
  if (type === 'card') {
    return (
      <div style={{ padding: '16px', display: 'flex', flexDirection: 'column', gap: '12px' }}>
        {Array.from({ length: rows }).map((_, i) => (
          <div key={i} className="skeleton" style={{ height: '60px', borderRadius: '6px' }} />
        ))}
      </div>
    )
  }

  return (
    <table style={{ width: '100%' }}>
      <tbody>
        {Array.from({ length: rows }).map((_, i) => (
          <tr key={i}>
            <td style={{ width: '20%' }}><div className="skeleton" style={{ height: '14px', width: '80%' }} /></td>
            <td style={{ width: '20%' }}><div className="skeleton" style={{ height: '14px', width: '70%' }} /></td>
            <td><div className="skeleton" style={{ height: '14px', width: '60%' }} /></td>
            <td><div className="skeleton" style={{ height: '14px', width: '50%' }} /></td>
            <td><div className="skeleton" style={{ height: '14px', width: '40%' }} /></td>
            <td><div className="skeleton" style={{ height: '14px', width: '60%' }} /></td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}

export { Skeleton }
