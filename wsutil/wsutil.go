/*
Package wsutil exports tools and helpers to simplify the work with WebSocket
protocol. The main purpose of this subpackage is to provide helpers that are
easy to use as possible.

Overview:

  // Read masked text message from peer and check utf8 encoding.
  header, err := ws.ReadHeader(conn)
  if err != nil {
	  // handle err
  }

  // Preapre to read payload.
  r := io.LimitReader(conn, header.Length)
  r = wsutil.NewCipherReader(r, header.Mask)
  r = wsutil.NewUTF8Reader(r)

  payload, err := ioutil.ReadAll(r)
  if err != nil {
	  // handle err
  }

You could get the same behavior using just `wsutil.Reader`:

  r := wsutil.Reader{
	  Source:    conn,
	  CheckUTF8: true,
  }

  payload, err := ioutil.ReadAll(r)
  if err != nil {
	  // handle err
  }

Or even simplest:

  payload, err := wsutil.ReadClientText(conn)
  if err != nil {
	  // handle err
  }

Pacakge provides also tool for buffered writing:

  // Create buffered writer, that will buffer output bytes and send them as
  // 128-length fragments (with exception on large writes, see the doc).
  writer := wsutil.NewWriterSize(conn, ws.StateServerSide, ws.OpText, 128)

  _, err := io.CopyN(writer, rand.Reader, 100)
  if err == nil {
	  err = writer.Flush()
  }
  if err != nil {
	  // handle error
  }

For more utils and helpers see the documentation.
*/
package wsutil
